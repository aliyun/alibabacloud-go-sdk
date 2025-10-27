// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachUserENIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachUserENIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachUserENIResponse
	GetStatusCode() *int32
	SetBody(v *AttachUserENIResponseBody) *AttachUserENIResponse
	GetBody() *AttachUserENIResponseBody
}

type AttachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachUserENIResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachUserENIResponse) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachUserENIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachUserENIResponse) GetBody() *AttachUserENIResponseBody {
	return s.Body
}

func (s *AttachUserENIResponse) SetHeaders(v map[string]*string) *AttachUserENIResponse {
	s.Headers = v
	return s
}

func (s *AttachUserENIResponse) SetStatusCode(v int32) *AttachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachUserENIResponse) SetBody(v *AttachUserENIResponseBody) *AttachUserENIResponse {
	s.Body = v
	return s
}

func (s *AttachUserENIResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
