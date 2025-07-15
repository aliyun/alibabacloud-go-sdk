// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMixStreamListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMixStreamListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMixStreamListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMixStreamListResponseBody) *DescribeMixStreamListResponse
	GetBody() *DescribeMixStreamListResponseBody
}

type DescribeMixStreamListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMixStreamListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMixStreamListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMixStreamListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMixStreamListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMixStreamListResponse) GetBody() *DescribeMixStreamListResponseBody {
	return s.Body
}

func (s *DescribeMixStreamListResponse) SetHeaders(v map[string]*string) *DescribeMixStreamListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMixStreamListResponse) SetStatusCode(v int32) *DescribeMixStreamListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMixStreamListResponse) SetBody(v *DescribeMixStreamListResponseBody) *DescribeMixStreamListResponse {
	s.Body = v
	return s
}

func (s *DescribeMixStreamListResponse) Validate() error {
	return dara.Validate(s)
}
