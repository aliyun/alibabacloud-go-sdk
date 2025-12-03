// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEaiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEaiResponse
	GetStatusCode() *int32
	SetBody(v *AttachEaiResponseBody) *AttachEaiResponse
	GetBody() *AttachEaiResponseBody
}

type AttachEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEaiResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEaiResponse) GoString() string {
	return s.String()
}

func (s *AttachEaiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEaiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEaiResponse) GetBody() *AttachEaiResponseBody {
	return s.Body
}

func (s *AttachEaiResponse) SetHeaders(v map[string]*string) *AttachEaiResponse {
	s.Headers = v
	return s
}

func (s *AttachEaiResponse) SetStatusCode(v int32) *AttachEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEaiResponse) SetBody(v *AttachEaiResponseBody) *AttachEaiResponse {
	s.Body = v
	return s
}

func (s *AttachEaiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
