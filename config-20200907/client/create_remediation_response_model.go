// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRemediationResponse
	GetStatusCode() *int32
	SetBody(v *CreateRemediationResponseBody) *CreateRemediationResponse
	GetBody() *CreateRemediationResponseBody
}

type CreateRemediationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRemediationResponse) GoString() string {
	return s.String()
}

func (s *CreateRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRemediationResponse) GetBody() *CreateRemediationResponseBody {
	return s.Body
}

func (s *CreateRemediationResponse) SetHeaders(v map[string]*string) *CreateRemediationResponse {
	s.Headers = v
	return s
}

func (s *CreateRemediationResponse) SetStatusCode(v int32) *CreateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemediationResponse) SetBody(v *CreateRemediationResponseBody) *CreateRemediationResponse {
	s.Body = v
	return s
}

func (s *CreateRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
