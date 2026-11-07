// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionDataResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionDataResponseBody) *QuerySessionDataResponse
	GetBody() *QuerySessionDataResponseBody
}

type QuerySessionDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionDataResponse) GetBody() *QuerySessionDataResponseBody {
	return s.Body
}

func (s *QuerySessionDataResponse) SetHeaders(v map[string]*string) *QuerySessionDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionDataResponse) SetStatusCode(v int32) *QuerySessionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionDataResponse) SetBody(v *QuerySessionDataResponseBody) *QuerySessionDataResponse {
	s.Body = v
	return s
}

func (s *QuerySessionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
