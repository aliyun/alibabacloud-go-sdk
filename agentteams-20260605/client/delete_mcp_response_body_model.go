// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMcpResponseBody
	GetSuccess() *bool
}

type DeleteMcpResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcpResponseBody) SetCode(v string) *DeleteMcpResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpResponseBody) SetMessage(v string) *DeleteMcpResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpResponseBody) SetRequestId(v string) *DeleteMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpResponseBody) SetSuccess(v bool) *DeleteMcpResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
