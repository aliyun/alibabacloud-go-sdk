// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMobilesCardSupportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMobilesCardSupportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMobilesCardSupportResponse
	GetStatusCode() *int32
	SetBody(v *CheckMobilesCardSupportResponseBody) *CheckMobilesCardSupportResponse
	GetBody() *CheckMobilesCardSupportResponseBody
}

type CheckMobilesCardSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMobilesCardSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMobilesCardSupportResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMobilesCardSupportResponse) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMobilesCardSupportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMobilesCardSupportResponse) GetBody() *CheckMobilesCardSupportResponseBody {
	return s.Body
}

func (s *CheckMobilesCardSupportResponse) SetHeaders(v map[string]*string) *CheckMobilesCardSupportResponse {
	s.Headers = v
	return s
}

func (s *CheckMobilesCardSupportResponse) SetStatusCode(v int32) *CheckMobilesCardSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMobilesCardSupportResponse) SetBody(v *CheckMobilesCardSupportResponseBody) *CheckMobilesCardSupportResponse {
	s.Body = v
	return s
}

func (s *CheckMobilesCardSupportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
