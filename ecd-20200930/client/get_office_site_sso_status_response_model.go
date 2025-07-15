// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficeSiteSsoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOfficeSiteSsoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOfficeSiteSsoStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetOfficeSiteSsoStatusResponseBody) *GetOfficeSiteSsoStatusResponse
	GetBody() *GetOfficeSiteSsoStatusResponseBody
}

type GetOfficeSiteSsoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfficeSiteSsoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOfficeSiteSsoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOfficeSiteSsoStatusResponse) GetBody() *GetOfficeSiteSsoStatusResponseBody {
	return s.Body
}

func (s *GetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *GetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) SetStatusCode(v int32) *GetOfficeSiteSsoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) SetBody(v *GetOfficeSiteSsoStatusResponseBody) *GetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) Validate() error {
	return dara.Validate(s)
}
