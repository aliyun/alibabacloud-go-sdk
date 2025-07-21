// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpfilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpfilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpfilterResponse
	GetStatusCode() *int32
	SetBody(v *AddIpfilterResponseBody) *AddIpfilterResponse
	GetBody() *AddIpfilterResponseBody
}

type AddIpfilterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpfilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpfilterResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpfilterResponse) GoString() string {
	return s.String()
}

func (s *AddIpfilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpfilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpfilterResponse) GetBody() *AddIpfilterResponseBody {
	return s.Body
}

func (s *AddIpfilterResponse) SetHeaders(v map[string]*string) *AddIpfilterResponse {
	s.Headers = v
	return s
}

func (s *AddIpfilterResponse) SetStatusCode(v int32) *AddIpfilterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpfilterResponse) SetBody(v *AddIpfilterResponseBody) *AddIpfilterResponse {
	s.Body = v
	return s
}

func (s *AddIpfilterResponse) Validate() error {
	return dara.Validate(s)
}
