// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceGroupNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceGroupNamesResponseBody) *DescribeDefenseResourceGroupNamesResponse
	GetBody() *DescribeDefenseResourceGroupNamesResponseBody
}

type DescribeDefenseResourceGroupNamesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceGroupNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceGroupNamesResponse) GetBody() *DescribeDefenseResourceGroupNamesResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetBody(v *DescribeDefenseResourceGroupNamesResponseBody) *DescribeDefenseResourceGroupNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
