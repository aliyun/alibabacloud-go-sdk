// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorksResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorksResponseBody) *QueryWorksResponse
	GetBody() *QueryWorksResponseBody
}

type QueryWorksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorksResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorksResponse) GetBody() *QueryWorksResponseBody {
	return s.Body
}

func (s *QueryWorksResponse) SetHeaders(v map[string]*string) *QueryWorksResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksResponse) SetStatusCode(v int32) *QueryWorksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksResponse) SetBody(v *QueryWorksResponseBody) *QueryWorksResponse {
	s.Body = v
	return s
}

func (s *QueryWorksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
