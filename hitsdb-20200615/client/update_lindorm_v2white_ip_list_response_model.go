// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2WhiteIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLindormV2WhiteIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLindormV2WhiteIpListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLindormV2WhiteIpListResponseBody) *UpdateLindormV2WhiteIpListResponse
	GetBody() *UpdateLindormV2WhiteIpListResponseBody
}

type UpdateLindormV2WhiteIpListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLindormV2WhiteIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLindormV2WhiteIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2WhiteIpListResponse) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2WhiteIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLindormV2WhiteIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLindormV2WhiteIpListResponse) GetBody() *UpdateLindormV2WhiteIpListResponseBody {
	return s.Body
}

func (s *UpdateLindormV2WhiteIpListResponse) SetHeaders(v map[string]*string) *UpdateLindormV2WhiteIpListResponse {
	s.Headers = v
	return s
}

func (s *UpdateLindormV2WhiteIpListResponse) SetStatusCode(v int32) *UpdateLindormV2WhiteIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLindormV2WhiteIpListResponse) SetBody(v *UpdateLindormV2WhiteIpListResponseBody) *UpdateLindormV2WhiteIpListResponse {
	s.Body = v
	return s
}

func (s *UpdateLindormV2WhiteIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
