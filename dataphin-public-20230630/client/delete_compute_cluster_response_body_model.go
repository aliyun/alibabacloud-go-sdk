// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteComputeClusterResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteComputeClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteComputeClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteComputeClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComputeClusterResponseBody
	GetSuccess() *bool
}

type DeleteComputeClusterResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComputeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteComputeClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteComputeClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteComputeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComputeClusterResponseBody) SetCode(v string) *DeleteComputeClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComputeClusterResponseBody) SetHttpStatusCode(v int32) *DeleteComputeClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteComputeClusterResponseBody) SetMessage(v string) *DeleteComputeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteComputeClusterResponseBody) SetRequestId(v string) *DeleteComputeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeClusterResponseBody) SetSuccess(v bool) *DeleteComputeClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComputeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
