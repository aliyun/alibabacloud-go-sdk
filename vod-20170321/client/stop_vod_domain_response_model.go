// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *StopVodDomainResponseBody) *StopVodDomainResponse
	GetBody() *StopVodDomainResponseBody
}

type StopVodDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StopVodDomainResponse) GoString() string {
	return s.String()
}

func (s *StopVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopVodDomainResponse) GetBody() *StopVodDomainResponseBody {
	return s.Body
}

func (s *StopVodDomainResponse) SetHeaders(v map[string]*string) *StopVodDomainResponse {
	s.Headers = v
	return s
}

func (s *StopVodDomainResponse) SetStatusCode(v int32) *StopVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopVodDomainResponse) SetBody(v *StopVodDomainResponseBody) *StopVodDomainResponse {
	s.Body = v
	return s
}

func (s *StopVodDomainResponse) Validate() error {
	return dara.Validate(s)
}
