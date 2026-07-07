// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteContactResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactResponseBody
	GetSuccess() *bool
}

type DeleteContactResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2ECA6FC9-7557-5576-AF5F-FC3E7BCC9C21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactResponseBody) SetCode(v int32) *DeleteContactResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactResponseBody) SetMessage(v string) *DeleteContactResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactResponseBody) SetSuccess(v bool) *DeleteContactResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactResponseBody) Validate() error {
	return dara.Validate(s)
}
