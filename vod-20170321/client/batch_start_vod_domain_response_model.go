// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartVodDomainResponseBody) *BatchStartVodDomainResponse
	GetBody() *BatchStartVodDomainResponseBody
}

type BatchStartVodDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartVodDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartVodDomainResponse) GetBody() *BatchStartVodDomainResponseBody {
	return s.Body
}

func (s *BatchStartVodDomainResponse) SetHeaders(v map[string]*string) *BatchStartVodDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartVodDomainResponse) SetStatusCode(v int32) *BatchStartVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartVodDomainResponse) SetBody(v *BatchStartVodDomainResponseBody) *BatchStartVodDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStartVodDomainResponse) Validate() error {
	return dara.Validate(s)
}
