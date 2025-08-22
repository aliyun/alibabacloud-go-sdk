// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopDcdnDomainResponseBody) *BatchStopDcdnDomainResponse
	GetBody() *BatchStopDcdnDomainResponseBody
}

type BatchStopDcdnDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopDcdnDomainResponse) GetBody() *BatchStopDcdnDomainResponseBody {
	return s.Body
}

func (s *BatchStopDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchStopDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopDcdnDomainResponse) SetStatusCode(v int32) *BatchStopDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopDcdnDomainResponse) SetBody(v *BatchStopDcdnDomainResponseBody) *BatchStopDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStopDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
