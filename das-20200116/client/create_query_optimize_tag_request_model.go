// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryOptimizeTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComments(v string) *CreateQueryOptimizeTagRequest
	GetComments() *string
	SetEngine(v string) *CreateQueryOptimizeTagRequest
	GetEngine() *string
	SetInstanceId(v string) *CreateQueryOptimizeTagRequest
	GetInstanceId() *string
	SetSqlIds(v string) *CreateQueryOptimizeTagRequest
	GetSqlIds() *string
	SetStatus(v int32) *CreateQueryOptimizeTagRequest
	GetStatus() *int32
	SetTags(v string) *CreateQueryOptimizeTagRequest
	GetTags() *string
}

type CreateQueryOptimizeTagRequest struct {
	// The remarks.
	//
	// The remarks can be 1 to 300 characters in length.
	//
	// example:
	//
	// Slow SQL queries of offline synchronization. No optimization is required.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL
	//
	// 	- **PolarDBMySQL**: PolarDB for MySQL
	//
	// 	- **PostgreSQL**: ApsaraDB RDS for PostgreSQL
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The SQL template IDs. You can call the [GetQueryOptimizeExecErrorStats](https://help.aliyun.com/document_detail/405261.html) operation to obtain the SQL template ID. Separate multiple SQL template IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 6068ce044e3dc9b903979672fb0b69df,d12515c015fc9f41a0778a9e1de0e941
	SqlIds *string `json:"SqlIds,omitempty" xml:"SqlIds,omitempty"`
	// The status of **Tags**. Valid values:
	//
	// 	- **0**: removes all tags added to the SQL templates that are specified by **SqlIds*	- and leaves **Tags*	- empty.
	//
	// 	- **1**: adds the tags specified by **Tags*	- to the SQL templates that are specified by **SqlIds**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The SQL tags. Separate multiple SQL tags with commas (,). Valid values:
	//
	// 	- **DAS_IMPORTANT**: The SQL template is important.
	//
	// 	- **DAS_NOT_IMPORTANT**: The SQL template is unimportant.
	//
	// 	- **USER_IGNORE**: The scheduling of the SQL template does not need to be optimized.
	//
	// 	- **DAS_IN_PLAN**: The scheduling of the SQL template needs to be optimized.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAS_IN_PLAN,DAS_NOT_IMPORTANT
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateQueryOptimizeTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryOptimizeTagRequest) GoString() string {
	return s.String()
}

func (s *CreateQueryOptimizeTagRequest) GetComments() *string {
	return s.Comments
}

func (s *CreateQueryOptimizeTagRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateQueryOptimizeTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQueryOptimizeTagRequest) GetSqlIds() *string {
	return s.SqlIds
}

func (s *CreateQueryOptimizeTagRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateQueryOptimizeTagRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateQueryOptimizeTagRequest) SetComments(v string) *CreateQueryOptimizeTagRequest {
	s.Comments = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) SetEngine(v string) *CreateQueryOptimizeTagRequest {
	s.Engine = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) SetInstanceId(v string) *CreateQueryOptimizeTagRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) SetSqlIds(v string) *CreateQueryOptimizeTagRequest {
	s.SqlIds = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) SetStatus(v int32) *CreateQueryOptimizeTagRequest {
	s.Status = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) SetTags(v string) *CreateQueryOptimizeTagRequest {
	s.Tags = &v
	return s
}

func (s *CreateQueryOptimizeTagRequest) Validate() error {
	return dara.Validate(s)
}
