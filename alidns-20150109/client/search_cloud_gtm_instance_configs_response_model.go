// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstanceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCloudGtmInstanceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCloudGtmInstanceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *SearchCloudGtmInstanceConfigsResponseBody) *SearchCloudGtmInstanceConfigsResponse
	GetBody() *SearchCloudGtmInstanceConfigsResponseBody
}

type SearchCloudGtmInstanceConfigsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCloudGtmInstanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCloudGtmInstanceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCloudGtmInstanceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCloudGtmInstanceConfigsResponse) GetBody() *SearchCloudGtmInstanceConfigsResponseBody {
	return s.Body
}

func (s *SearchCloudGtmInstanceConfigsResponse) SetHeaders(v map[string]*string) *SearchCloudGtmInstanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponse) SetStatusCode(v int32) *SearchCloudGtmInstanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponse) SetBody(v *SearchCloudGtmInstanceConfigsResponseBody) *SearchCloudGtmInstanceConfigsResponse {
	s.Body = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
