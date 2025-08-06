// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVodDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetVodDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetVodDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetVodDomainConfigsResponseBody) *BatchSetVodDomainConfigsResponse
	GetBody() *BatchSetVodDomainConfigsResponseBody
}

type BatchSetVodDomainConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetVodDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetVodDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVodDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetVodDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetVodDomainConfigsResponse) GetBody() *BatchSetVodDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetVodDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetVodDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetVodDomainConfigsResponse) SetStatusCode(v int32) *BatchSetVodDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetVodDomainConfigsResponse) SetBody(v *BatchSetVodDomainConfigsResponseBody) *BatchSetVodDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetVodDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
