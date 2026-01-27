// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostInfoByDbsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostInfoByDbsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostInfoByDbsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostInfoByDbsInstanceResponseBody) *DescribeCostInfoByDbsInstanceResponse
	GetBody() *DescribeCostInfoByDbsInstanceResponseBody
}

type DescribeCostInfoByDbsInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostInfoByDbsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostInfoByDbsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostInfoByDbsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostInfoByDbsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostInfoByDbsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostInfoByDbsInstanceResponse) GetBody() *DescribeCostInfoByDbsInstanceResponseBody {
	return s.Body
}

func (s *DescribeCostInfoByDbsInstanceResponse) SetHeaders(v map[string]*string) *DescribeCostInfoByDbsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponse) SetStatusCode(v int32) *DescribeCostInfoByDbsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponse) SetBody(v *DescribeCostInfoByDbsInstanceResponseBody) *DescribeCostInfoByDbsInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
