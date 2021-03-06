package testutilRepo

import (
	commonrepo "github.com/agriardyan/common-lib-v2/repo"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockDBOperation struct {
	mock.Mock
	defaultCollection string
}

func (dbOp *MockDBOperation) InsertOne(data interface{}) (*mongo.InsertOneResult, error) {
	return dbOp.InsertOneAtColl(dbOp.defaultCollection, data)
}

func (dbOp *MockDBOperation) InsertMany(data []interface{}) (*mongo.InsertManyResult, error) {
	return dbOp.InsertManyAtColl(dbOp.defaultCollection, data)
}

func (dbOp *MockDBOperation) FindOne(filter, impl interface{}) (interface{}, error) {
	return dbOp.FindOneAtColl(dbOp.defaultCollection, filter, impl)
}

func (dbOp *MockDBOperation) FindOneAndUpdate(filter, update interface{}) (*mongo.UpdateResult, error) {
	return dbOp.FindOneAndUpdateAtColl(dbOp.defaultCollection, filter, update)
}

func (dbOp *MockDBOperation) FindOneAndUpsert(filter, update interface{}) (*mongo.UpdateResult, error) {
	return dbOp.FindOneAndUpdateAtColl(dbOp.defaultCollection, filter, update)
}

func (dbOp *MockDBOperation) Find(filter, impl interface{}) (interface{}, error) {
	return dbOp.FindAtColl(dbOp.defaultCollection, filter, impl)
}

func (dbOp *MockDBOperation) FindPagedSorted(pagingSortingReq commonrepo.PagingSortingRequest, filter, impl interface{}) (interface{}, *mongopagination.PaginationData, error) {
	return dbOp.FindAtCollPagedSorted(dbOp.defaultCollection, pagingSortingReq, filter, impl)
}

func (dbOp *MockDBOperation) Count(filter interface{}) (int64, error) {
	return dbOp.CountAtColl(dbOp.defaultCollection, filter)
}

func (dbOp *MockDBOperation) IsExist(filter interface{}) (bool, error) {
	return dbOp.IsExistAtColl(dbOp.defaultCollection, filter)
}

func (dbOp *MockDBOperation) InsertOneAtColl(collection string, data interface{}) (*mongo.InsertOneResult, error) {
	args := dbOp.Called(collection, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (dbOp *MockDBOperation) InsertManyAtColl(collection string, data []interface{}) (*mongo.InsertManyResult, error) {
	args := dbOp.Called(collection, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertManyResult), args.Error(1)
}

func (dbOp *MockDBOperation) FindOneAtColl(collection string, filter, impl interface{}) (interface{}, error) {
	args := dbOp.Called(collection, filter, impl)
	return args.Get(0), args.Error(1)
}

func (dbOp *MockDBOperation) FindOneAndUpdateAtColl(collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	args := dbOp.Called(collection, filter, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (dbOp *MockDBOperation) FindOneAndUpsertAtColl(collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	args := dbOp.Called(collection, filter, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (dbOp *MockDBOperation) FindAtColl(collection string, filter, impl interface{}) (interface{}, error) {
	args := dbOp.Called(collection, filter, impl)
	return args.Get(0), args.Error(1)
}

func (dbOp *MockDBOperation) FindAtCollPagedSorted(collection string, pagingSortingReq commonrepo.PagingSortingRequest, filter, impl interface{}) (interface{}, *mongopagination.PaginationData, error) {
	args := dbOp.Called(collection, pagingSortingReq, filter, impl)
	return args.Get(0), args.Get(1).(*mongopagination.PaginationData), args.Error(2)
}

func (dbOp *MockDBOperation) CountAtColl(collection string, filter interface{}) (int64, error) {
	args := dbOp.Called(collection, filter)
	return int64(args.Int(0)), args.Error(1)
}

func (dbOp *MockDBOperation) IsExistAtColl(collection string, filter interface{}) (bool, error) {
	args := dbOp.Called(collection, filter)
	return args.Bool(0), args.Error(1)
}
