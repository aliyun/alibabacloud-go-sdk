// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchFileResponse
	GetStatusCode() *int32
	SetBody(v *FetchFileResponseBody) *FetchFileResponse
	GetBody() *FetchFileResponseBody
}

type FetchFileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchFileResponse) GoString() string {
	return s.String()
}

func (s *FetchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchFileResponse) GetBody() *FetchFileResponseBody {
	return s.Body
}

func (s *FetchFileResponse) SetHeaders(v map[string]*string) *FetchFileResponse {
	s.Headers = v
	return s
}

func (s *FetchFileResponse) SetStatusCode(v int32) *FetchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchFileResponse) SetBody(v *FetchFileResponseBody) *FetchFileResponse {
	s.Body = v
	return s
}

func (s *FetchFileResponse) Validate() error {
	return dara.Validate(s)
}
