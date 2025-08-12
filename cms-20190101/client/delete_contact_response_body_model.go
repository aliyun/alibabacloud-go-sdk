// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContactResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactResponseBody
	GetSuccess() *bool
}

type DeleteContactResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 50D4CFE1-01E5-4543-939C-18BC01E3EC1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
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

func (s *DeleteContactResponseBody) GetCode() *string {
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

func (s *DeleteContactResponseBody) SetCode(v string) *DeleteContactResponseBody {
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
