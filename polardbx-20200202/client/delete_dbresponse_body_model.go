// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteDBResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDBResponseBody
	GetSuccess() *bool
}

type DeleteDBResponseBody struct {
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDBResponseBody) SetMessage(v string) *DeleteDBResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDBResponseBody) SetRequestId(v string) *DeleteDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBResponseBody) SetSuccess(v bool) *DeleteDBResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDBResponseBody) Validate() error {
	return dara.Validate(s)
}
