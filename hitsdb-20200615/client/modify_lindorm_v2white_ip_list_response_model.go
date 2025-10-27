// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2WhiteIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLindormV2WhiteIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLindormV2WhiteIpListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLindormV2WhiteIpListResponseBody) *ModifyLindormV2WhiteIpListResponse
	GetBody() *ModifyLindormV2WhiteIpListResponseBody
}

type ModifyLindormV2WhiteIpListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLindormV2WhiteIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLindormV2WhiteIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLindormV2WhiteIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLindormV2WhiteIpListResponse) GetBody() *ModifyLindormV2WhiteIpListResponseBody {
	return s.Body
}

func (s *ModifyLindormV2WhiteIpListResponse) SetHeaders(v map[string]*string) *ModifyLindormV2WhiteIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponse) SetStatusCode(v int32) *ModifyLindormV2WhiteIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponse) SetBody(v *ModifyLindormV2WhiteIpListResponseBody) *ModifyLindormV2WhiteIpListResponse {
	s.Body = v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
