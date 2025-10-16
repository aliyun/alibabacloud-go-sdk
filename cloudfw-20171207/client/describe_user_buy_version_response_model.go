// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBuyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBuyVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBuyVersionResponseBody) *DescribeUserBuyVersionResponse
	GetBody() *DescribeUserBuyVersionResponseBody
}

type DescribeUserBuyVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBuyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBuyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBuyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBuyVersionResponse) GetBody() *DescribeUserBuyVersionResponseBody {
	return s.Body
}

func (s *DescribeUserBuyVersionResponse) SetHeaders(v map[string]*string) *DescribeUserBuyVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBuyVersionResponse) SetStatusCode(v int32) *DescribeUserBuyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBuyVersionResponse) SetBody(v *DescribeUserBuyVersionResponseBody) *DescribeUserBuyVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBuyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
