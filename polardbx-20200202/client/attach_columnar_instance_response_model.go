// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachColumnarInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachColumnarInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachColumnarInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AttachColumnarInstanceResponseBody) *AttachColumnarInstanceResponse
	GetBody() *AttachColumnarInstanceResponseBody
}

type AttachColumnarInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachColumnarInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachColumnarInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachColumnarInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachColumnarInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachColumnarInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachColumnarInstanceResponse) GetBody() *AttachColumnarInstanceResponseBody {
	return s.Body
}

func (s *AttachColumnarInstanceResponse) SetHeaders(v map[string]*string) *AttachColumnarInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachColumnarInstanceResponse) SetStatusCode(v int32) *AttachColumnarInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachColumnarInstanceResponse) SetBody(v *AttachColumnarInstanceResponseBody) *AttachColumnarInstanceResponse {
	s.Body = v
	return s
}

func (s *AttachColumnarInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
