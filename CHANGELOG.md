# Changelog

## v0.0.1
### Added
- New err type unique constraint violation
- Common auditing struct
- New constructor and default collection name for DBOperation
- Common repo db operation: FindAtColl, FindOneAtColl, InsertAtColl, UpdateAtColl to enable user specify custom collection
- Suppport for pagination and sort in common repo
- Parser and validator for Basic authentication
- Common repo db operation: Find, FindOne, Insert, Update
- Common response across services such as UnknownError or DBProblem response
- Common error code and explanation across services
- Swagger html error message generator
- Check whether error is mongodb's duplication write error