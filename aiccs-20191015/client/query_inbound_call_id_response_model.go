// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInboundCallIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInboundCallIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInboundCallIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryInboundCallIdResponseBody) *QueryInboundCallIdResponse
	GetBody() *QueryInboundCallIdResponseBody
}

type QueryInboundCallIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInboundCallIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInboundCallIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInboundCallIdResponse) GoString() string {
	return s.String()
}

func (s *QueryInboundCallIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInboundCallIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInboundCallIdResponse) GetBody() *QueryInboundCallIdResponseBody {
	return s.Body
}

func (s *QueryInboundCallIdResponse) SetHeaders(v map[string]*string) *QueryInboundCallIdResponse {
	s.Headers = v
	return s
}

func (s *QueryInboundCallIdResponse) SetStatusCode(v int32) *QueryInboundCallIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInboundCallIdResponse) SetBody(v *QueryInboundCallIdResponseBody) *QueryInboundCallIdResponse {
	s.Body = v
	return s
}

func (s *QueryInboundCallIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
