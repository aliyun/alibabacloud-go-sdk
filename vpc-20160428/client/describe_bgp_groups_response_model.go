// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBgpGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBgpGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBgpGroupsResponseBody) *DescribeBgpGroupsResponse
	GetBody() *DescribeBgpGroupsResponseBody
}

type DescribeBgpGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBgpGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBgpGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBgpGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBgpGroupsResponse) GetBody() *DescribeBgpGroupsResponseBody {
	return s.Body
}

func (s *DescribeBgpGroupsResponse) SetHeaders(v map[string]*string) *DescribeBgpGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpGroupsResponse) SetStatusCode(v int32) *DescribeBgpGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpGroupsResponse) SetBody(v *DescribeBgpGroupsResponseBody) *DescribeBgpGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeBgpGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
