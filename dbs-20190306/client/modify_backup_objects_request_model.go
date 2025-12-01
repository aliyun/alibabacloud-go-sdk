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
	// The backup objects that are defined in a JSON string in the following format:
	//
	//     [
	//
	//         {
	//
	//             "DBName":"The name of the database that you want to back up",
	//
	//             "SchemaName":"The name of the schema that you want to back up",
	//
	//             "TableIncludes":[{
	//
	//             	"TableName":"The name of the table that you want to back up"
	//
	//             }],
	//
	//             "TableExcludes":[{
	//
	//                 "TableName":"The name of the table that you want to exclude during the backup"
	//
	//             }]
	//
	//         }
	//
	//     ]
	//
	// 	- If you specify only `DBName` and do not specify objects of lower levels, all objects in the database are backed up.
	//
	// 	- If you specify `DBName` and some objects of lower levels, only the specified objects are backed up by default. You can use the following regular expressions to define object names:
	//
	//     	- A period `.` matches any single character except `\\r\\n`.
	//
	//     	- An asterisk `*` matches zero or more occurrences of a preceding subexpression. For example, `h.*llo` matches strings such as `hllo` and `heeeello`.
	//
	//     	- A question mark `?` matches zero or one occurrence of a preceding subexpression. For example, `h.?llo` matches strings such as `hllo` and `hello`, but not `haello`.
	//
	//     	- Character set `[Characters]` matches a character included in the brackets ([ ]). For example, `h[ae]llo` matches `hallo` and `hello`.
	//
	//     	- Negative character set `[^Characters]` does not match a character in the brackets ([ ]). For example, `h[^ae]llo` matches `hcllo` and `hdllo`, but not `hallo` or `hello`.
	//
	//     	- Character range `[character1-character2]` matches any character included in the range from `character1 to character2`, such as `[0-9]` and `[a-z]`.
	//
	// >  `SchemaName` and `NewSchemaName` apply only to SQL Server databases. Use `DBName` and `NewDBName` to specify the names of other databases.
	//
	// This parameter is required.
	//
	// example:
	//
	// [     {         "DBName":"17xxx92xxxx374",         "SchemaName":"gy9xnj8xxxxxx",         "TableIncludes":[{         	"TableName":"test1"         }],         "TableExcludes":[{             "TableName":"test2"         }]     } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
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
