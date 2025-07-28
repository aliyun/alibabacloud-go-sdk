// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveStatisticResponseBody) *DescribeSensitiveStatisticResponse
	GetBody() *DescribeSensitiveStatisticResponseBody
}

type DescribeSensitiveStatisticResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveStatisticResponse) GetBody() *DescribeSensitiveStatisticResponseBody {
	return s.Body
}

func (s *DescribeSensitiveStatisticResponse) SetHeaders(v map[string]*string) *DescribeSensitiveStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveStatisticResponse) SetStatusCode(v int32) *DescribeSensitiveStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveStatisticResponse) SetBody(v *DescribeSensitiveStatisticResponseBody) *DescribeSensitiveStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveStatisticResponse) Validate() error {
	return dara.Validate(s)
}
