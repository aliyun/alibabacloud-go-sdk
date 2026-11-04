// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshGbSubDeviceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshGbSubDeviceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshGbSubDeviceListResponse
	GetStatusCode() *int32
	SetBody(v *RefreshGbSubDeviceListResponseBody) *RefreshGbSubDeviceListResponse
	GetBody() *RefreshGbSubDeviceListResponseBody
}

type RefreshGbSubDeviceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshGbSubDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshGbSubDeviceListResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshGbSubDeviceListResponse) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshGbSubDeviceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshGbSubDeviceListResponse) GetBody() *RefreshGbSubDeviceListResponseBody {
	return s.Body
}

func (s *RefreshGbSubDeviceListResponse) SetHeaders(v map[string]*string) *RefreshGbSubDeviceListResponse {
	s.Headers = v
	return s
}

func (s *RefreshGbSubDeviceListResponse) SetStatusCode(v int32) *RefreshGbSubDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshGbSubDeviceListResponse) SetBody(v *RefreshGbSubDeviceListResponseBody) *RefreshGbSubDeviceListResponse {
	s.Body = v
	return s
}

func (s *RefreshGbSubDeviceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
