// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopVodDomainResponseBody) *BatchStopVodDomainResponse
	GetBody() *BatchStopVodDomainResponseBody
}

type BatchStopVodDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopVodDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopVodDomainResponse) GetBody() *BatchStopVodDomainResponseBody {
	return s.Body
}

func (s *BatchStopVodDomainResponse) SetHeaders(v map[string]*string) *BatchStopVodDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopVodDomainResponse) SetStatusCode(v int32) *BatchStopVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopVodDomainResponse) SetBody(v *BatchStopVodDomainResponseBody) *BatchStopVodDomainResponse {
	s.Body = v
	return s
}

func (s *BatchStopVodDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
