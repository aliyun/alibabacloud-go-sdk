// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveVerifyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveVerifyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveVerifyContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveVerifyContentResponseBody) *DescribeLiveVerifyContentResponse
	GetBody() *DescribeLiveVerifyContentResponseBody
}

type DescribeLiveVerifyContentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveVerifyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveVerifyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveVerifyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveVerifyContentResponse) GetBody() *DescribeLiveVerifyContentResponseBody {
	return s.Body
}

func (s *DescribeLiveVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeLiveVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveVerifyContentResponse) SetStatusCode(v int32) *DescribeLiveVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveVerifyContentResponse) SetBody(v *DescribeLiveVerifyContentResponseBody) *DescribeLiveVerifyContentResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveVerifyContentResponse) Validate() error {
	return dara.Validate(s)
}
