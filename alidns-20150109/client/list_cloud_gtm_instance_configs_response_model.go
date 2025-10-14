// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstanceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmInstanceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmInstanceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmInstanceConfigsResponseBody) *ListCloudGtmInstanceConfigsResponse
	GetBody() *ListCloudGtmInstanceConfigsResponseBody
}

type ListCloudGtmInstanceConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmInstanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmInstanceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmInstanceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmInstanceConfigsResponse) GetBody() *ListCloudGtmInstanceConfigsResponseBody {
	return s.Body
}

func (s *ListCloudGtmInstanceConfigsResponse) SetHeaders(v map[string]*string) *ListCloudGtmInstanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponse) SetStatusCode(v int32) *ListCloudGtmInstanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponse) SetBody(v *ListCloudGtmInstanceConfigsResponseBody) *ListCloudGtmInstanceConfigsResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
