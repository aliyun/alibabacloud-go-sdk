// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitUploadPolicyResponseBody) *CreateTransitUploadPolicyResponse
	GetBody() *CreateTransitUploadPolicyResponseBody
}

type CreateTransitUploadPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitUploadPolicyResponse) GetBody() *CreateTransitUploadPolicyResponseBody {
	return s.Body
}

func (s *CreateTransitUploadPolicyResponse) SetHeaders(v map[string]*string) *CreateTransitUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitUploadPolicyResponse) SetStatusCode(v int32) *CreateTransitUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitUploadPolicyResponse) SetBody(v *CreateTransitUploadPolicyResponseBody) *CreateTransitUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateTransitUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
