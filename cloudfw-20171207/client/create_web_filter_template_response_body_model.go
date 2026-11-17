// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWebFilterTemplateResponseBody
	GetRequestId() *string
	SetWebFilterTemplateId(v string) *CreateWebFilterTemplateResponseBody
	GetWebFilterTemplateId() *string
}

type CreateWebFilterTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the newly created web filtering template.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s CreateWebFilterTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebFilterTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebFilterTemplateResponseBody) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *CreateWebFilterTemplateResponseBody) SetRequestId(v string) *CreateWebFilterTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebFilterTemplateResponseBody) SetWebFilterTemplateId(v string) *CreateWebFilterTemplateResponseBody {
	s.WebFilterTemplateId = &v
	return s
}

func (s *CreateWebFilterTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
