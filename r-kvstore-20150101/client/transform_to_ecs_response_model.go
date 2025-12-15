// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToEcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformToEcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformToEcsResponse
	GetStatusCode() *int32
	SetBody(v *TransformToEcsResponseBody) *TransformToEcsResponse
	GetBody() *TransformToEcsResponseBody
}

type TransformToEcsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformToEcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformToEcsResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformToEcsResponse) GoString() string {
	return s.String()
}

func (s *TransformToEcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformToEcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformToEcsResponse) GetBody() *TransformToEcsResponseBody {
	return s.Body
}

func (s *TransformToEcsResponse) SetHeaders(v map[string]*string) *TransformToEcsResponse {
	s.Headers = v
	return s
}

func (s *TransformToEcsResponse) SetStatusCode(v int32) *TransformToEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformToEcsResponse) SetBody(v *TransformToEcsResponseBody) *TransformToEcsResponse {
	s.Body = v
	return s
}

func (s *TransformToEcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
