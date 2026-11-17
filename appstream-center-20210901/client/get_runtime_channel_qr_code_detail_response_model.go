// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelQrCodeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuntimeChannelQrCodeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuntimeChannelQrCodeDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetRuntimeChannelQrCodeDetailResponseBody) *GetRuntimeChannelQrCodeDetailResponse
	GetBody() *GetRuntimeChannelQrCodeDetailResponseBody
}

type GetRuntimeChannelQrCodeDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuntimeChannelQrCodeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuntimeChannelQrCodeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelQrCodeDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelQrCodeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuntimeChannelQrCodeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuntimeChannelQrCodeDetailResponse) GetBody() *GetRuntimeChannelQrCodeDetailResponseBody {
	return s.Body
}

func (s *GetRuntimeChannelQrCodeDetailResponse) SetHeaders(v map[string]*string) *GetRuntimeChannelQrCodeDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponse) SetStatusCode(v int32) *GetRuntimeChannelQrCodeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponse) SetBody(v *GetRuntimeChannelQrCodeDetailResponseBody) *GetRuntimeChannelQrCodeDetailResponse {
	s.Body = v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
