// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartCdnDomainResponseBody) *BatchStartCdnDomainResponse
	GetBody() *BatchStartCdnDomainResponseBody
}

type BatchStartCdnDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartCdnDomainResponse) GetBody() *BatchStartCdnDomainResponseBody {
	return s.Body
}

func (s *BatchStartCdnDomainResponse) SetHeaders(v map[string]*string) *BatchStartCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartCdnDomainResponse) SetStatusCode(v int32) *BatchStartCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartCdnDomainResponse) SetBody(v *BatchStartCdnDomainResponseBody) *BatchStartCdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStartCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
