// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEssayCorrectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunEssayCorrectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunEssayCorrectionResponse
	GetStatusCode() *int32
	SetBody(v *RunEssayCorrectionResponseBody) *RunEssayCorrectionResponse
	GetBody() *RunEssayCorrectionResponseBody
}

type RunEssayCorrectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunEssayCorrectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunEssayCorrectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponse) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunEssayCorrectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunEssayCorrectionResponse) GetBody() *RunEssayCorrectionResponseBody {
	return s.Body
}

func (s *RunEssayCorrectionResponse) SetHeaders(v map[string]*string) *RunEssayCorrectionResponse {
	s.Headers = v
	return s
}

func (s *RunEssayCorrectionResponse) SetStatusCode(v int32) *RunEssayCorrectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunEssayCorrectionResponse) SetBody(v *RunEssayCorrectionResponseBody) *RunEssayCorrectionResponse {
	s.Body = v
	return s
}

func (s *RunEssayCorrectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
