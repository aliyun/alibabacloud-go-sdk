// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterScenesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterScenesResponseBody) *DescribeCasterScenesResponse
	GetBody() *DescribeCasterScenesResponseBody
}

type DescribeCasterScenesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterScenesResponse) GetBody() *DescribeCasterScenesResponseBody {
	return s.Body
}

func (s *DescribeCasterScenesResponse) SetHeaders(v map[string]*string) *DescribeCasterScenesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterScenesResponse) SetStatusCode(v int32) *DescribeCasterScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterScenesResponse) SetBody(v *DescribeCasterScenesResponseBody) *DescribeCasterScenesResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterScenesResponse) Validate() error {
	return dara.Validate(s)
}
