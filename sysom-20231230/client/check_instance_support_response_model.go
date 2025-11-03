// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceSupportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceSupportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceSupportResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceSupportResponseBody) *CheckInstanceSupportResponse
	GetBody() *CheckInstanceSupportResponseBody
}

type CheckInstanceSupportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceSupportResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceSupportResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceSupportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceSupportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceSupportResponse) GetBody() *CheckInstanceSupportResponseBody {
	return s.Body
}

func (s *CheckInstanceSupportResponse) SetHeaders(v map[string]*string) *CheckInstanceSupportResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceSupportResponse) SetStatusCode(v int32) *CheckInstanceSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceSupportResponse) SetBody(v *CheckInstanceSupportResponseBody) *CheckInstanceSupportResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceSupportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
