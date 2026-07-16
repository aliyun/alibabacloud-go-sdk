// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMatchSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMatchSiteResponse
	GetStatusCode() *int32
	SetBody(v *GetMatchSiteResponseBody) *GetMatchSiteResponse
	GetBody() *GetMatchSiteResponseBody
}

type GetMatchSiteResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMatchSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMatchSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMatchSiteResponse) GoString() string {
	return s.String()
}

func (s *GetMatchSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMatchSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMatchSiteResponse) GetBody() *GetMatchSiteResponseBody {
	return s.Body
}

func (s *GetMatchSiteResponse) SetHeaders(v map[string]*string) *GetMatchSiteResponse {
	s.Headers = v
	return s
}

func (s *GetMatchSiteResponse) SetStatusCode(v int32) *GetMatchSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMatchSiteResponse) SetBody(v *GetMatchSiteResponseBody) *GetMatchSiteResponse {
	s.Body = v
	return s
}

func (s *GetMatchSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
