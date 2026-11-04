// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGbSubDeviceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGbSubDeviceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGbSubDeviceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryGbSubDeviceListResponseBody) *QueryGbSubDeviceListResponse
	GetBody() *QueryGbSubDeviceListResponseBody
}

type QueryGbSubDeviceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGbSubDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGbSubDeviceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGbSubDeviceListResponse) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGbSubDeviceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGbSubDeviceListResponse) GetBody() *QueryGbSubDeviceListResponseBody {
	return s.Body
}

func (s *QueryGbSubDeviceListResponse) SetHeaders(v map[string]*string) *QueryGbSubDeviceListResponse {
	s.Headers = v
	return s
}

func (s *QueryGbSubDeviceListResponse) SetStatusCode(v int32) *QueryGbSubDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGbSubDeviceListResponse) SetBody(v *QueryGbSubDeviceListResponseBody) *QueryGbSubDeviceListResponse {
	s.Body = v
	return s
}

func (s *QueryGbSubDeviceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
