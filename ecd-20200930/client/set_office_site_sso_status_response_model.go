// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOfficeSiteSsoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetOfficeSiteSsoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetOfficeSiteSsoStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetOfficeSiteSsoStatusResponseBody) *SetOfficeSiteSsoStatusResponse
	GetBody() *SetOfficeSiteSsoStatusResponseBody
}

type SetOfficeSiteSsoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetOfficeSiteSsoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetOfficeSiteSsoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetOfficeSiteSsoStatusResponse) GetBody() *SetOfficeSiteSsoStatusResponseBody {
	return s.Body
}

func (s *SetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *SetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetOfficeSiteSsoStatusResponse) SetStatusCode(v int32) *SetOfficeSiteSsoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetOfficeSiteSsoStatusResponse) SetBody(v *SetOfficeSiteSsoStatusResponseBody) *SetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

func (s *SetOfficeSiteSsoStatusResponse) Validate() error {
	return dara.Validate(s)
}
