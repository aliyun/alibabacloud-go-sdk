// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *GetQueryOptimizeTagRequest
	GetEngine() *string
	SetInstanceId(v string) *GetQueryOptimizeTagRequest
	GetInstanceId() *string
	SetSqlId(v string) *GetQueryOptimizeTagRequest
	GetSqlId() *string
}

type GetQueryOptimizeTagRequest struct {
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
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The SQL template ID. You can call the [GetQueryOptimizeDataStats](https://help.aliyun.com/document_detail/405261.html) operation to query the SQL template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29d9fef63e347d39c3436658a5fe5f2b
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetQueryOptimizeTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeTagRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeTagRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeTagRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeTagRequest) SetEngine(v string) *GetQueryOptimizeTagRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeTagRequest) SetInstanceId(v string) *GetQueryOptimizeTagRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeTagRequest) SetSqlId(v string) *GetQueryOptimizeTagRequest {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeTagRequest) Validate() error {
	return dara.Validate(s)
}
