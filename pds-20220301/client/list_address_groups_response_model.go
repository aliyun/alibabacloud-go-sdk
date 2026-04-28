// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddressGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddressGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAddressGroupsResponseBody) *ListAddressGroupsResponse
	GetBody() *ListAddressGroupsResponseBody
}

type ListAddressGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddressGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddressGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddressGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddressGroupsResponse) GetBody() *ListAddressGroupsResponseBody {
	return s.Body
}

func (s *ListAddressGroupsResponse) SetHeaders(v map[string]*string) *ListAddressGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAddressGroupsResponse) SetStatusCode(v int32) *ListAddressGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddressGroupsResponse) SetBody(v *ListAddressGroupsResponseBody) *ListAddressGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAddressGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
