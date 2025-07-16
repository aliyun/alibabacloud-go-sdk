// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopCdnDomainResponseBody) *BatchStopCdnDomainResponse
	GetBody() *BatchStopCdnDomainResponseBody
}

type BatchStopCdnDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopCdnDomainResponse) GetBody() *BatchStopCdnDomainResponseBody {
	return s.Body
}

func (s *BatchStopCdnDomainResponse) SetHeaders(v map[string]*string) *BatchStopCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopCdnDomainResponse) SetStatusCode(v int32) *BatchStopCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopCdnDomainResponse) SetBody(v *BatchStopCdnDomainResponseBody) *BatchStopCdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStopCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
