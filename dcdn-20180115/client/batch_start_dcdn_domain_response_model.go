// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartDcdnDomainResponseBody) *BatchStartDcdnDomainResponse
	GetBody() *BatchStartDcdnDomainResponseBody
}

type BatchStartDcdnDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartDcdnDomainResponse) GetBody() *BatchStartDcdnDomainResponseBody {
	return s.Body
}

func (s *BatchStartDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchStartDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartDcdnDomainResponse) SetStatusCode(v int32) *BatchStartDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartDcdnDomainResponse) SetBody(v *BatchStartDcdnDomainResponseBody) *BatchStartDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStartDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
