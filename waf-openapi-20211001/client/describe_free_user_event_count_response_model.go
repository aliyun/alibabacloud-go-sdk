// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFreeUserEventCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFreeUserEventCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFreeUserEventCountResponseBody) *DescribeFreeUserEventCountResponse
	GetBody() *DescribeFreeUserEventCountResponseBody
}

type DescribeFreeUserEventCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFreeUserEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFreeUserEventCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFreeUserEventCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFreeUserEventCountResponse) GetBody() *DescribeFreeUserEventCountResponseBody {
	return s.Body
}

func (s *DescribeFreeUserEventCountResponse) SetHeaders(v map[string]*string) *DescribeFreeUserEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeFreeUserEventCountResponse) SetStatusCode(v int32) *DescribeFreeUserEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFreeUserEventCountResponse) SetBody(v *DescribeFreeUserEventCountResponseBody) *DescribeFreeUserEventCountResponse {
	s.Body = v
	return s
}

func (s *DescribeFreeUserEventCountResponse) Validate() error {
	return dara.Validate(s)
}
