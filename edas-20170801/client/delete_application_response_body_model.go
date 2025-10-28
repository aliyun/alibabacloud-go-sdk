// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DeleteApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *DeleteApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApplicationResponseBody
	GetRequestId() *string
}

type DeleteApplicationResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// 0b8e3c0b-5818-430*************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34DFE9FDV****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationResponseBody) SetChangeOrderId(v string) *DeleteApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetCode(v int32) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
