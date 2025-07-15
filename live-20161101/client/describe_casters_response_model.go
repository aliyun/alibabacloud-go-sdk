// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCastersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCastersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCastersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCastersResponseBody) *DescribeCastersResponse
	GetBody() *DescribeCastersResponseBody
}

type DescribeCastersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCastersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCastersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponse) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCastersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCastersResponse) GetBody() *DescribeCastersResponseBody {
	return s.Body
}

func (s *DescribeCastersResponse) SetHeaders(v map[string]*string) *DescribeCastersResponse {
	s.Headers = v
	return s
}

func (s *DescribeCastersResponse) SetStatusCode(v int32) *DescribeCastersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCastersResponse) SetBody(v *DescribeCastersResponseBody) *DescribeCastersResponse {
	s.Body = v
	return s
}

func (s *DescribeCastersResponse) Validate() error {
	return dara.Validate(s)
}
