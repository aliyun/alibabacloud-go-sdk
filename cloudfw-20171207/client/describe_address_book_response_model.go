// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddressBookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddressBookResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddressBookResponseBody) *DescribeAddressBookResponse
	GetBody() *DescribeAddressBookResponseBody
}

type DescribeAddressBookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddressBookResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddressBookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddressBookResponse) GetBody() *DescribeAddressBookResponseBody {
	return s.Body
}

func (s *DescribeAddressBookResponse) SetHeaders(v map[string]*string) *DescribeAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressBookResponse) SetStatusCode(v int32) *DescribeAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressBookResponse) SetBody(v *DescribeAddressBookResponseBody) *DescribeAddressBookResponse {
	s.Body = v
	return s
}

func (s *DescribeAddressBookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
