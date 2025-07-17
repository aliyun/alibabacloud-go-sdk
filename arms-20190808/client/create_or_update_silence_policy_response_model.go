// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSilencePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateSilencePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateSilencePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateSilencePolicyResponseBody) *CreateOrUpdateSilencePolicyResponse
	GetBody() *CreateOrUpdateSilencePolicyResponseBody
}

type CreateOrUpdateSilencePolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateSilencePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateSilencePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateSilencePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateSilencePolicyResponse) GetBody() *CreateOrUpdateSilencePolicyResponseBody {
	return s.Body
}

func (s *CreateOrUpdateSilencePolicyResponse) SetHeaders(v map[string]*string) *CreateOrUpdateSilencePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponse) SetStatusCode(v int32) *CreateOrUpdateSilencePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponse) SetBody(v *CreateOrUpdateSilencePolicyResponseBody) *CreateOrUpdateSilencePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponse) Validate() error {
	return dara.Validate(s)
}
