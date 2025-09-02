// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeIOResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListNodeIOResponseBodyData) *ListNodeIOResponseBody
	GetData() []*ListNodeIOResponseBodyData
	SetErrorCode(v string) *ListNodeIOResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListNodeIOResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListNodeIOResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListNodeIOResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodeIOResponseBody
	GetSuccess() *bool
}

type ListNodeIOResponseBody struct {
	// The node information.
	Data []*ListNodeIOResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ProjectNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The project does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeIOResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeIOResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeIOResponseBody) GetData() []*ListNodeIOResponseBodyData {
	return s.Data
}

func (s *ListNodeIOResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNodeIOResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListNodeIOResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodeIOResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeIOResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodeIOResponseBody) SetData(v []*ListNodeIOResponseBodyData) *ListNodeIOResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeIOResponseBody) SetErrorCode(v string) *ListNodeIOResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNodeIOResponseBody) SetErrorMessage(v string) *ListNodeIOResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListNodeIOResponseBody) SetHttpStatusCode(v int32) *ListNodeIOResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodeIOResponseBody) SetRequestId(v string) *ListNodeIOResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeIOResponseBody) SetSuccess(v bool) *ListNodeIOResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodeIOResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodeIOResponseBodyData struct {
	// The name of the ancestor or descendant node.
	//
	// example:
	//
	// dataworks_a.1234_out
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 123123
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the input or output table.
	//
	// example:
	//
	// dataworks_a.datastudio_tenant_waitres_alarm
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListNodeIOResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNodeIOResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeIOResponseBodyData) GetData() *string {
	return s.Data
}

func (s *ListNodeIOResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodeIOResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *ListNodeIOResponseBodyData) SetData(v string) *ListNodeIOResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListNodeIOResponseBodyData) SetNodeId(v int64) *ListNodeIOResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *ListNodeIOResponseBodyData) SetTableName(v string) *ListNodeIOResponseBodyData {
	s.TableName = &v
	return s
}

func (s *ListNodeIOResponseBodyData) Validate() error {
	return dara.Validate(s)
}
