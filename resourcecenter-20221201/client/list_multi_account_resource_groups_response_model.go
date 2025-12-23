// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiAccountResourceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiAccountResourceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiAccountResourceGroupsResponseBody) *ListMultiAccountResourceGroupsResponse
	GetBody() *ListMultiAccountResourceGroupsResponseBody
}

type ListMultiAccountResourceGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiAccountResourceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiAccountResourceGroupsResponse) GetBody() *ListMultiAccountResourceGroupsResponseBody {
	return s.Body
}

func (s *ListMultiAccountResourceGroupsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetStatusCode(v int32) *ListMultiAccountResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetBody(v *ListMultiAccountResourceGroupsResponseBody) *ListMultiAccountResourceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
