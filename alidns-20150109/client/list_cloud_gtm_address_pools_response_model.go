// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmAddressPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmAddressPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmAddressPoolsResponseBody) *ListCloudGtmAddressPoolsResponse
	GetBody() *ListCloudGtmAddressPoolsResponseBody
}

type ListCloudGtmAddressPoolsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmAddressPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmAddressPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmAddressPoolsResponse) GetBody() *ListCloudGtmAddressPoolsResponseBody {
	return s.Body
}

func (s *ListCloudGtmAddressPoolsResponse) SetHeaders(v map[string]*string) *ListCloudGtmAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponse) SetStatusCode(v int32) *ListCloudGtmAddressPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponse) SetBody(v *ListCloudGtmAddressPoolsResponseBody) *ListCloudGtmAddressPoolsResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
