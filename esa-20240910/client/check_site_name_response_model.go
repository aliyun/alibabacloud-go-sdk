// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSiteNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSiteNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckSiteNameResponseBody) *CheckSiteNameResponse
	GetBody() *CheckSiteNameResponseBody
}

type CheckSiteNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSiteNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSiteNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteNameResponse) GoString() string {
	return s.String()
}

func (s *CheckSiteNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSiteNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSiteNameResponse) GetBody() *CheckSiteNameResponseBody {
	return s.Body
}

func (s *CheckSiteNameResponse) SetHeaders(v map[string]*string) *CheckSiteNameResponse {
	s.Headers = v
	return s
}

func (s *CheckSiteNameResponse) SetStatusCode(v int32) *CheckSiteNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSiteNameResponse) SetBody(v *CheckSiteNameResponseBody) *CheckSiteNameResponse {
	s.Body = v
	return s
}

func (s *CheckSiteNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
