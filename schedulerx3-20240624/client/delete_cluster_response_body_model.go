// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteClusterResponseBody
	GetSuccess() *bool
}

type DeleteClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F131C3E0-3FAA-5FA4-A6F3-E974D69EF3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteClusterResponseBody) SetCode(v int32) *DeleteClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
