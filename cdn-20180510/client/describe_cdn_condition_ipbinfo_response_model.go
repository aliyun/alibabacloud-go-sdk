// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnConditionIPBInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnConditionIPBInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnConditionIPBInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnConditionIPBInfoResponseBody) *DescribeCdnConditionIPBInfoResponse
	GetBody() *DescribeCdnConditionIPBInfoResponseBody
}

type DescribeCdnConditionIPBInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnConditionIPBInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnConditionIPBInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnConditionIPBInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnConditionIPBInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnConditionIPBInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnConditionIPBInfoResponse) GetBody() *DescribeCdnConditionIPBInfoResponseBody {
	return s.Body
}

func (s *DescribeCdnConditionIPBInfoResponse) SetHeaders(v map[string]*string) *DescribeCdnConditionIPBInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponse) SetStatusCode(v int32) *DescribeCdnConditionIPBInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponse) SetBody(v *DescribeCdnConditionIPBInfoResponseBody) *DescribeCdnConditionIPBInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponse) Validate() error {
	return dara.Validate(s)
}
