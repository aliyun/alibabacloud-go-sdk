// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCode(v string) *CreateCommandResponseBody
	GetDomainCode() *string
	SetRequestId(v string) *CreateCommandResponseBody
	GetRequestId() *string
	SetToolId(v string) *CreateCommandResponseBody
	GetToolId() *string
}

type CreateCommandResponseBody struct {
	// example:
	//
	// 72893434
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3435676586
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
}

func (s CreateCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommandResponseBody) GetDomainCode() *string {
	return s.DomainCode
}

func (s *CreateCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommandResponseBody) GetToolId() *string {
	return s.ToolId
}

func (s *CreateCommandResponseBody) SetDomainCode(v string) *CreateCommandResponseBody {
	s.DomainCode = &v
	return s
}

func (s *CreateCommandResponseBody) SetRequestId(v string) *CreateCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommandResponseBody) SetToolId(v string) *CreateCommandResponseBody {
	s.ToolId = &v
	return s
}

func (s *CreateCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
