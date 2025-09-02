// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateClusterConfigsResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateClusterConfigsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateClusterConfigsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateClusterConfigsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateClusterConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterConfigsResponseBody
	GetSuccess() *bool
}

type UpdateClusterConfigsResponseBody struct {
	// Indicates whether the update was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 101011005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid.Cluster.ClusterNotFound
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. It is used to locate logs and troubleshoot problems.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterConfigsResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateClusterConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateClusterConfigsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateClusterConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateClusterConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterConfigsResponseBody) SetData(v bool) *UpdateClusterConfigsResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) SetErrorCode(v string) *UpdateClusterConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) SetErrorMessage(v string) *UpdateClusterConfigsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) SetHttpStatusCode(v int32) *UpdateClusterConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) SetRequestId(v string) *UpdateClusterConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) SetSuccess(v bool) *UpdateClusterConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
