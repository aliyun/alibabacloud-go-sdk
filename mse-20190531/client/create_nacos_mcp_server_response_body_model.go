// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateNacosMcpServerResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateNacosMcpServerResponseBody
	GetRequestId() *string
}

type CreateNacosMcpServerResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AF21683A-29C7-4853-AC0F-B5ADEE4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNacosMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNacosMcpServerResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateNacosMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNacosMcpServerResponseBody) SetData(v bool) *CreateNacosMcpServerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateNacosMcpServerResponseBody) SetRequestId(v string) *CreateNacosMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNacosMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
