// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxgGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAxgGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAxgGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAxgGroupResponseBody
	GetRequestId() *string
}

type DeleteAxgGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9297B722-A016-43FB-B51A-E54050D9369D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAxgGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAxgGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAxgGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAxgGroupResponseBody) SetCode(v string) *DeleteAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAxgGroupResponseBody) SetMessage(v string) *DeleteAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAxgGroupResponseBody) SetRequestId(v string) *DeleteAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAxgGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
