// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupObjects(v string) *ModifyBackupObjectsRequest
	GetBackupObjects() *string
	SetBackupPlanId(v string) *ModifyBackupObjectsRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *ModifyBackupObjectsRequest
	GetClientToken() *string
	SetOwnerId(v string) *ModifyBackupObjectsRequest
	GetOwnerId() *string
}

type ModifyBackupObjectsRequest struct {
	// The backup objects, specified as a JSON string. The structure is as follows:
	//
	// ```
	//
	// [
	//
	//     {
	//
	//         "DBName":"The name of the database to back up",
	//
	//         "SchemaName":"The name of the schema to back up",
	//
	//         "TableIncludes":[{
	//
	//         	"TableName":"The name of the table to back up"
	//
	//         }],
	//
	//         "TableExcludes":[{
	//
	//             "TableName":"The name of a table in the database that you do not want to back up"
	//
	//         }]
	//
	//     }
	//
	// ]
	//
	// ```
	//
	// - If you specify only `DBName` without configuring rules for sub-objects, all objects in the database are backed up.
	//
	// - If you specify `DBName` and configure rules for some objects, any objects without configured rules are not backed up by default. The following regular expressions are supported for defining object names:
	//
	//   - A period (`.`) matches any single character except `
	//
	// `.
	//
	//   - An asterisk (`*`) matches the preceding subexpression zero or more times. For example, `h*llo` matches `hllo` and `heeeello`.
	//
	//   - A question mark (`?`) matches the preceding subexpression zero or one time. For example, `h.?llo` matches `hllo` and `hello`, but not `haello`.
	//
	//   - A character set `[characters]` matches any single character within the brackets. For example, `h[aello]` matches `hallo` and `hello`.
	//
	//   - A negated character set `[^characters]` matches any single character not within the brackets. For example, `h[^ae]llo` matches `hcllo` and `hdllo`, but not `hallo` or `hello`.
	//
	//   - A character range `[character1-character2]` matches any character within the specified range, such as `[0-9]` or `[a-z]`.
	//
	// > `SchemaName` and `NewSchemaName` are used only for SQL Server. For other database engines, use `DBName` and `NewDBName` to specify database names.
	//
	// This parameter is required.
	//
	// example:
	//
	// [     {         "DBName":"17xxx92xxxx374",         "SchemaName":"gy9xnj8xxxxxx",         "TableIncludes":[{         	"TableName":"test1"         }],         "TableExcludes":[{             "TableName":"test2"         }]     } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A client token to ensure the request is idempotent. This prevents the same request from being submitted multiple times.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyBackupObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupObjectsRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsRequest) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *ModifyBackupObjectsRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupObjectsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupObjectsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyBackupObjectsRequest) SetBackupObjects(v string) *ModifyBackupObjectsRequest {
	s.BackupObjects = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetBackupPlanId(v string) *ModifyBackupObjectsRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetClientToken(v string) *ModifyBackupObjectsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetOwnerId(v string) *ModifyBackupObjectsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupObjectsRequest) Validate() error {
	return dara.Validate(s)
}
