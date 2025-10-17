// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDBNameZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDBNameZonalResponse
	GetStatusCode() *int32
	SetBody(v *CheckDBNameZonalResponseBody) *CheckDBNameZonalResponse
	GetBody() *CheckDBNameZonalResponseBody
}

type CheckDBNameZonalResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDBNameZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDBNameZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameZonalResponse) GoString() string {
	return s.String()
}

func (s *CheckDBNameZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDBNameZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDBNameZonalResponse) GetBody() *CheckDBNameZonalResponseBody {
	return s.Body
}

func (s *CheckDBNameZonalResponse) SetHeaders(v map[string]*string) *CheckDBNameZonalResponse {
	s.Headers = v
	return s
}

func (s *CheckDBNameZonalResponse) SetStatusCode(v int32) *CheckDBNameZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDBNameZonalResponse) SetBody(v *CheckDBNameZonalResponseBody) *CheckDBNameZonalResponse {
	s.Body = v
	return s
}

func (s *CheckDBNameZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
