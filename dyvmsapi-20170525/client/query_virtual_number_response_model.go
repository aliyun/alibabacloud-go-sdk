// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVirtualNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVirtualNumberResponse
	GetStatusCode() *int32
	SetBody(v *QueryVirtualNumberResponseBody) *QueryVirtualNumberResponse
	GetBody() *QueryVirtualNumberResponseBody
}

type QueryVirtualNumberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVirtualNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVirtualNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberResponse) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVirtualNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVirtualNumberResponse) GetBody() *QueryVirtualNumberResponseBody {
	return s.Body
}

func (s *QueryVirtualNumberResponse) SetHeaders(v map[string]*string) *QueryVirtualNumberResponse {
	s.Headers = v
	return s
}

func (s *QueryVirtualNumberResponse) SetStatusCode(v int32) *QueryVirtualNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVirtualNumberResponse) SetBody(v *QueryVirtualNumberResponseBody) *QueryVirtualNumberResponse {
	s.Body = v
	return s
}

func (s *QueryVirtualNumberResponse) Validate() error {
	return dara.Validate(s)
}
