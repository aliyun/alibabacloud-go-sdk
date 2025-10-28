// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDeployGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ChangeDeployGroupResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ChangeDeployGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *ChangeDeployGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeDeployGroupResponseBody
	GetRequestId() *string
}

type ChangeDeployGroupResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// 435f-regfr4********************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F9E4-FDS4-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeDeployGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ChangeDeployGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ChangeDeployGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeDeployGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDeployGroupResponseBody) SetChangeOrderId(v string) *ChangeDeployGroupResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetCode(v int32) *ChangeDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetMessage(v string) *ChangeDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetRequestId(v string) *ChangeDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
