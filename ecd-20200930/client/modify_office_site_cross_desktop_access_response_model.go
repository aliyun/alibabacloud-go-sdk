// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteCrossDesktopAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteCrossDesktopAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteCrossDesktopAccessResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteCrossDesktopAccessResponseBody) *ModifyOfficeSiteCrossDesktopAccessResponse
	GetBody() *ModifyOfficeSiteCrossDesktopAccessResponseBody
}

type ModifyOfficeSiteCrossDesktopAccessResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteCrossDesktopAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) GetBody() *ModifyOfficeSiteCrossDesktopAccessResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetStatusCode(v int32) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetBody(v *ModifyOfficeSiteCrossDesktopAccessResponseBody) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
