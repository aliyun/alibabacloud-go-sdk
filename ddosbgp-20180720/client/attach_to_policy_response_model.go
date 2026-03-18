// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachToPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachToPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachToPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AttachToPolicyResponseBody) *AttachToPolicyResponse
	GetBody() *AttachToPolicyResponseBody
}

type AttachToPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachToPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachToPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachToPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachToPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachToPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachToPolicyResponse) GetBody() *AttachToPolicyResponseBody {
	return s.Body
}

func (s *AttachToPolicyResponse) SetHeaders(v map[string]*string) *AttachToPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachToPolicyResponse) SetStatusCode(v int32) *AttachToPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachToPolicyResponse) SetBody(v *AttachToPolicyResponseBody) *AttachToPolicyResponse {
	s.Body = v
	return s
}

func (s *AttachToPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
