// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGraphResponseBody
	GetCode() *string
	SetDataSourceId(v int64) *CreateGraphResponseBody
	GetDataSourceId() *int64
	SetGraphName(v string) *CreateGraphResponseBody
	GetGraphName() *string
	SetMessage(v string) *CreateGraphResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGraphResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *CreateGraphResponseBody
	GetSchemaVersion() *string
	SetSyncStatus(v string) *CreateGraphResponseBody
	GetSyncStatus() *string
}

type CreateGraphResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 198001
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// The graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The version.
	//
	// example:
	//
	// 0.0.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The OSS synchronization and extraction status. Valid values:
	//
	// - -1: Not set (default).
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// SUCCESS
	SyncStatus *string `json:"syncStatus,omitempty" xml:"syncStatus,omitempty"`
}

func (s CreateGraphResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGraphResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGraphResponseBody) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateGraphResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *CreateGraphResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGraphResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGraphResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *CreateGraphResponseBody) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *CreateGraphResponseBody) SetCode(v string) *CreateGraphResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGraphResponseBody) SetDataSourceId(v int64) *CreateGraphResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateGraphResponseBody) SetGraphName(v string) *CreateGraphResponseBody {
	s.GraphName = &v
	return s
}

func (s *CreateGraphResponseBody) SetMessage(v string) *CreateGraphResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGraphResponseBody) SetRequestId(v string) *CreateGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGraphResponseBody) SetSchemaVersion(v string) *CreateGraphResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *CreateGraphResponseBody) SetSyncStatus(v string) *CreateGraphResponseBody {
	s.SyncStatus = &v
	return s
}

func (s *CreateGraphResponseBody) Validate() error {
	return dara.Validate(s)
}
