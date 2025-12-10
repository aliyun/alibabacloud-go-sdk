// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsWithAuthDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupsWithAuthDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupsWithAuthDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupsWithAuthDetailsResponseBody) *ListResourceGroupsWithAuthDetailsResponse
	GetBody() *ListResourceGroupsWithAuthDetailsResponseBody
}

type ListResourceGroupsWithAuthDetailsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsWithAuthDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupsWithAuthDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupsWithAuthDetailsResponse) GetBody() *ListResourceGroupsWithAuthDetailsResponseBody {
	return s.Body
}

func (s *ListResourceGroupsWithAuthDetailsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsWithAuthDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponse) SetStatusCode(v int32) *ListResourceGroupsWithAuthDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponse) SetBody(v *ListResourceGroupsWithAuthDetailsResponseBody) *ListResourceGroupsWithAuthDetailsResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
