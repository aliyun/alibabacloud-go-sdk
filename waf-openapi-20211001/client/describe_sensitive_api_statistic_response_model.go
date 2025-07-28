// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveApiStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveApiStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveApiStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveApiStatisticResponseBody) *DescribeSensitiveApiStatisticResponse
	GetBody() *DescribeSensitiveApiStatisticResponseBody
}

type DescribeSensitiveApiStatisticResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveApiStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveApiStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveApiStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveApiStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveApiStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveApiStatisticResponse) GetBody() *DescribeSensitiveApiStatisticResponseBody {
	return s.Body
}

func (s *DescribeSensitiveApiStatisticResponse) SetHeaders(v map[string]*string) *DescribeSensitiveApiStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveApiStatisticResponse) SetStatusCode(v int32) *DescribeSensitiveApiStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveApiStatisticResponse) SetBody(v *DescribeSensitiveApiStatisticResponseBody) *DescribeSensitiveApiStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveApiStatisticResponse) Validate() error {
	return dara.Validate(s)
}
