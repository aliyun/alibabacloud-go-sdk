// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRegistrantProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRegistrantProfilesResponse
	GetStatusCode() *int32
	SetBody(v *QueryRegistrantProfilesResponseBody) *QueryRegistrantProfilesResponse
	GetBody() *QueryRegistrantProfilesResponseBody
}

type QueryRegistrantProfilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRegistrantProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRegistrantProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponse) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRegistrantProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRegistrantProfilesResponse) GetBody() *QueryRegistrantProfilesResponseBody {
	return s.Body
}

func (s *QueryRegistrantProfilesResponse) SetHeaders(v map[string]*string) *QueryRegistrantProfilesResponse {
	s.Headers = v
	return s
}

func (s *QueryRegistrantProfilesResponse) SetStatusCode(v int32) *QueryRegistrantProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegistrantProfilesResponse) SetBody(v *QueryRegistrantProfilesResponseBody) *QueryRegistrantProfilesResponse {
	s.Body = v
	return s
}

func (s *QueryRegistrantProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
