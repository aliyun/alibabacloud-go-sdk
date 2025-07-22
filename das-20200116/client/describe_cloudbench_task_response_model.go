// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudbenchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudbenchTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudbenchTaskResponseBody) *DescribeCloudbenchTaskResponse
	GetBody() *DescribeCloudbenchTaskResponseBody
}

type DescribeCloudbenchTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudbenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudbenchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudbenchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudbenchTaskResponse) GetBody() *DescribeCloudbenchTaskResponseBody {
	return s.Body
}

func (s *DescribeCloudbenchTaskResponse) SetHeaders(v map[string]*string) *DescribeCloudbenchTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudbenchTaskResponse) SetStatusCode(v int32) *DescribeCloudbenchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudbenchTaskResponse) SetBody(v *DescribeCloudbenchTaskResponseBody) *DescribeCloudbenchTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudbenchTaskResponse) Validate() error {
	return dara.Validate(s)
}
