// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSystemPropertyTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSystemPropertyTemplatesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSystemPropertyTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSystemPropertyTemplatesResponseBody
	GetRequestId() *string
}

type DeleteSystemPropertyTemplatesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSystemPropertyTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSystemPropertyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSystemPropertyTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSystemPropertyTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSystemPropertyTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSystemPropertyTemplatesResponseBody) SetCode(v string) *DeleteSystemPropertyTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponseBody) SetMessage(v string) *DeleteSystemPropertyTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponseBody) SetRequestId(v string) *DeleteSystemPropertyTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
