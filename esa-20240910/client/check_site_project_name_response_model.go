// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteProjectNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSiteProjectNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSiteProjectNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckSiteProjectNameResponseBody) *CheckSiteProjectNameResponse
	GetBody() *CheckSiteProjectNameResponseBody
}

type CheckSiteProjectNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSiteProjectNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSiteProjectNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteProjectNameResponse) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSiteProjectNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSiteProjectNameResponse) GetBody() *CheckSiteProjectNameResponseBody {
	return s.Body
}

func (s *CheckSiteProjectNameResponse) SetHeaders(v map[string]*string) *CheckSiteProjectNameResponse {
	s.Headers = v
	return s
}

func (s *CheckSiteProjectNameResponse) SetStatusCode(v int32) *CheckSiteProjectNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSiteProjectNameResponse) SetBody(v *CheckSiteProjectNameResponseBody) *CheckSiteProjectNameResponse {
	s.Body = v
	return s
}

func (s *CheckSiteProjectNameResponse) Validate() error {
	return dara.Validate(s)
}
