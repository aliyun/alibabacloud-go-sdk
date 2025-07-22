// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *GetQueryOptimizeExecErrorSampleRequest
	GetEngine() *string
	SetInstanceId(v string) *GetQueryOptimizeExecErrorSampleRequest
	GetInstanceId() *string
	SetSqlId(v string) *GetQueryOptimizeExecErrorSampleRequest
	GetSqlId() *string
	SetTime(v string) *GetQueryOptimizeExecErrorSampleRequest
	GetTime() *string
}

type GetQueryOptimizeExecErrorSampleRequest struct {
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PostgreSQL**
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
	// The SQL template ID. You can call the [GetQueryOptimizeExecErrorStats](https://help.aliyun.com/document_detail/405235.html) operation to obtain the SQL template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The date to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642953600000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleRequest) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetQueryOptimizeExecErrorSampleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeExecErrorSampleRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeExecErrorSampleRequest) GetTime() *string {
	return s.Time
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetEngine(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.Engine = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetInstanceId(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetSqlId(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) SetTime(v string) *GetQueryOptimizeExecErrorSampleRequest {
	s.Time = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleRequest) Validate() error {
	return dara.Validate(s)
}
