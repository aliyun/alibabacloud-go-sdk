// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCloudGtmAddressPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCloudGtmAddressPoolsResponse
	GetStatusCode() *int32
	SetBody(v *SearchCloudGtmAddressPoolsResponseBody) *SearchCloudGtmAddressPoolsResponse
	GetBody() *SearchCloudGtmAddressPoolsResponseBody
}

type SearchCloudGtmAddressPoolsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCloudGtmAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCloudGtmAddressPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCloudGtmAddressPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCloudGtmAddressPoolsResponse) GetBody() *SearchCloudGtmAddressPoolsResponseBody {
	return s.Body
}

func (s *SearchCloudGtmAddressPoolsResponse) SetHeaders(v map[string]*string) *SearchCloudGtmAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponse) SetStatusCode(v int32) *SearchCloudGtmAddressPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponse) SetBody(v *SearchCloudGtmAddressPoolsResponseBody) *SearchCloudGtmAddressPoolsResponse {
	s.Body = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
