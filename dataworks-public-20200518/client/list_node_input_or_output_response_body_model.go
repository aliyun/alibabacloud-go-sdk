// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInputOrOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListNodeInputOrOutputResponseBodyData) *ListNodeInputOrOutputResponseBody
	GetData() []*ListNodeInputOrOutputResponseBodyData
	SetErrorCode(v string) *ListNodeInputOrOutputResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListNodeInputOrOutputResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListNodeInputOrOutputResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListNodeInputOrOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodeInputOrOutputResponseBody
	GetSuccess() *bool
}

type ListNodeInputOrOutputResponseBody struct {
	// The ancestor or descendant nodes.
	Data []*ListNodeInputOrOutputResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListNodeInputOrOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInputOrOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInputOrOutputResponseBody) GetData() []*ListNodeInputOrOutputResponseBodyData {
	return s.Data
}

func (s *ListNodeInputOrOutputResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNodeInputOrOutputResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListNodeInputOrOutputResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodeInputOrOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeInputOrOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodeInputOrOutputResponseBody) SetData(v []*ListNodeInputOrOutputResponseBodyData) *ListNodeInputOrOutputResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) SetErrorCode(v string) *ListNodeInputOrOutputResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) SetErrorMessage(v string) *ListNodeInputOrOutputResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) SetHttpStatusCode(v int32) *ListNodeInputOrOutputResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) SetRequestId(v string) *ListNodeInputOrOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) SetSuccess(v bool) *ListNodeInputOrOutputResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodeInputOrOutputResponseBodyData struct {
	// The name of the ancestor or descendant node.
	//
	// example:
	//
	// xxxx.123141254_out
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 1234667
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter does not take effect. You cannot obtain the parameter settings.
	//
	// example:
	//
	// dwd_xxx_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListNodeInputOrOutputResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInputOrOutputResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeInputOrOutputResponseBodyData) GetData() *string {
	return s.Data
}

func (s *ListNodeInputOrOutputResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodeInputOrOutputResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *ListNodeInputOrOutputResponseBodyData) SetData(v string) *ListNodeInputOrOutputResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBodyData) SetNodeId(v int64) *ListNodeInputOrOutputResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBodyData) SetTableName(v string) *ListNodeInputOrOutputResponseBodyData {
	s.TableName = &v
	return s
}

func (s *ListNodeInputOrOutputResponseBodyData) Validate() error {
	return dara.Validate(s)
}
