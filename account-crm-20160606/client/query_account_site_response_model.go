// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountSiteResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountSiteResponseBody) *QueryAccountSiteResponse
	GetBody() *QueryAccountSiteResponseBody
}

type QueryAccountSiteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSiteResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountSiteResponse) GetBody() *QueryAccountSiteResponseBody {
	return s.Body
}

func (s *QueryAccountSiteResponse) SetHeaders(v map[string]*string) *QueryAccountSiteResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountSiteResponse) SetStatusCode(v int32) *QueryAccountSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountSiteResponse) SetBody(v *QueryAccountSiteResponseBody) *QueryAccountSiteResponse {
	s.Body = v
	return s
}

func (s *QueryAccountSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
