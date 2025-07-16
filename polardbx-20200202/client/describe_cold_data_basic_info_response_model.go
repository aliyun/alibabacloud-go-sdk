// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdDataBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColdDataBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColdDataBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColdDataBasicInfoResponseBody) *DescribeColdDataBasicInfoResponse
	GetBody() *DescribeColdDataBasicInfoResponseBody
}

type DescribeColdDataBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdDataBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdDataBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColdDataBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColdDataBasicInfoResponse) GetBody() *DescribeColdDataBasicInfoResponseBody {
	return s.Body
}

func (s *DescribeColdDataBasicInfoResponse) SetHeaders(v map[string]*string) *DescribeColdDataBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdDataBasicInfoResponse) SetStatusCode(v int32) *DescribeColdDataBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponse) SetBody(v *DescribeColdDataBasicInfoResponseBody) *DescribeColdDataBasicInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeColdDataBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
