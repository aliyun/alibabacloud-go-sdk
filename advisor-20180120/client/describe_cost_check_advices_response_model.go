// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckAdvicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostCheckAdvicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostCheckAdvicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostCheckAdvicesResponseBody) *DescribeCostCheckAdvicesResponse
	GetBody() *DescribeCostCheckAdvicesResponseBody
}

type DescribeCostCheckAdvicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostCheckAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostCheckAdvicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostCheckAdvicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostCheckAdvicesResponse) GetBody() *DescribeCostCheckAdvicesResponseBody {
	return s.Body
}

func (s *DescribeCostCheckAdvicesResponse) SetHeaders(v map[string]*string) *DescribeCostCheckAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostCheckAdvicesResponse) SetStatusCode(v int32) *DescribeCostCheckAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponse) SetBody(v *DescribeCostCheckAdvicesResponseBody) *DescribeCostCheckAdvicesResponse {
	s.Body = v
	return s
}

func (s *DescribeCostCheckAdvicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
