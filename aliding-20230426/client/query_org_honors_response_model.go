// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrgHonorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrgHonorsResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrgHonorsResponseBody) *QueryOrgHonorsResponse
	GetBody() *QueryOrgHonorsResponseBody
}

type QueryOrgHonorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgHonorsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrgHonorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrgHonorsResponse) GetBody() *QueryOrgHonorsResponseBody {
	return s.Body
}

func (s *QueryOrgHonorsResponse) SetHeaders(v map[string]*string) *QueryOrgHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgHonorsResponse) SetStatusCode(v int32) *QueryOrgHonorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgHonorsResponse) SetBody(v *QueryOrgHonorsResponseBody) *QueryOrgHonorsResponse {
	s.Body = v
	return s
}

func (s *QueryOrgHonorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
