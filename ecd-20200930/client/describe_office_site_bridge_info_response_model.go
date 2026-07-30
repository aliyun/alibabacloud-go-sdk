// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSiteBridgeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOfficeSiteBridgeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOfficeSiteBridgeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOfficeSiteBridgeInfoResponseBody) *DescribeOfficeSiteBridgeInfoResponse
	GetBody() *DescribeOfficeSiteBridgeInfoResponseBody
}

type DescribeOfficeSiteBridgeInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfficeSiteBridgeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOfficeSiteBridgeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSiteBridgeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSiteBridgeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOfficeSiteBridgeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOfficeSiteBridgeInfoResponse) GetBody() *DescribeOfficeSiteBridgeInfoResponseBody {
	return s.Body
}

func (s *DescribeOfficeSiteBridgeInfoResponse) SetHeaders(v map[string]*string) *DescribeOfficeSiteBridgeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponse) SetStatusCode(v int32) *DescribeOfficeSiteBridgeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponse) SetBody(v *DescribeOfficeSiteBridgeInfoResponseBody) *DescribeOfficeSiteBridgeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
