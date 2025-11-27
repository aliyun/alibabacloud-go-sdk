// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmsQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmsQualificationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmsQualificationResponseBody) *DeleteSmsQualificationResponse
	GetBody() *DeleteSmsQualificationResponseBody
}

type DeleteSmsQualificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmsQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmsQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsQualificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmsQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmsQualificationResponse) GetBody() *DeleteSmsQualificationResponseBody {
	return s.Body
}

func (s *DeleteSmsQualificationResponse) SetHeaders(v map[string]*string) *DeleteSmsQualificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsQualificationResponse) SetStatusCode(v int32) *DeleteSmsQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmsQualificationResponse) SetBody(v *DeleteSmsQualificationResponseBody) *DeleteSmsQualificationResponse {
	s.Body = v
	return s
}

func (s *DeleteSmsQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
