// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenHttpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenHttpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenHttpsResponse
	GetStatusCode() *int32
	SetBody(v *OpenHttpsResponseBody) *OpenHttpsResponse
	GetBody() *OpenHttpsResponseBody
}

type OpenHttpsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenHttpsResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenHttpsResponse) GoString() string {
	return s.String()
}

func (s *OpenHttpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenHttpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenHttpsResponse) GetBody() *OpenHttpsResponseBody {
	return s.Body
}

func (s *OpenHttpsResponse) SetHeaders(v map[string]*string) *OpenHttpsResponse {
	s.Headers = v
	return s
}

func (s *OpenHttpsResponse) SetStatusCode(v int32) *OpenHttpsResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenHttpsResponse) SetBody(v *OpenHttpsResponseBody) *OpenHttpsResponse {
	s.Body = v
	return s
}

func (s *OpenHttpsResponse) Validate() error {
	return dara.Validate(s)
}
