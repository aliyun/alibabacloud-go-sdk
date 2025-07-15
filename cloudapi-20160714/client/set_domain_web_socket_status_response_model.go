// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainWebSocketStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainWebSocketStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainWebSocketStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainWebSocketStatusResponseBody) *SetDomainWebSocketStatusResponse
	GetBody() *SetDomainWebSocketStatusResponseBody
}

type SetDomainWebSocketStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainWebSocketStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainWebSocketStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainWebSocketStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainWebSocketStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainWebSocketStatusResponse) GetBody() *SetDomainWebSocketStatusResponseBody {
	return s.Body
}

func (s *SetDomainWebSocketStatusResponse) SetHeaders(v map[string]*string) *SetDomainWebSocketStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainWebSocketStatusResponse) SetStatusCode(v int32) *SetDomainWebSocketStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainWebSocketStatusResponse) SetBody(v *SetDomainWebSocketStatusResponseBody) *SetDomainWebSocketStatusResponse {
	s.Body = v
	return s
}

func (s *SetDomainWebSocketStatusResponse) Validate() error {
	return dara.Validate(s)
}
