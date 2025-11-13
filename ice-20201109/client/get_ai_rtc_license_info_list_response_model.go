// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcLicenseInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiRtcLicenseInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiRtcLicenseInfoListResponse
	GetStatusCode() *int32
	SetBody(v *GetAiRtcLicenseInfoListResponseBody) *GetAiRtcLicenseInfoListResponse
	GetBody() *GetAiRtcLicenseInfoListResponseBody
}

type GetAiRtcLicenseInfoListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiRtcLicenseInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiRtcLicenseInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcLicenseInfoListResponse) GoString() string {
	return s.String()
}

func (s *GetAiRtcLicenseInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiRtcLicenseInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiRtcLicenseInfoListResponse) GetBody() *GetAiRtcLicenseInfoListResponseBody {
	return s.Body
}

func (s *GetAiRtcLicenseInfoListResponse) SetHeaders(v map[string]*string) *GetAiRtcLicenseInfoListResponse {
	s.Headers = v
	return s
}

func (s *GetAiRtcLicenseInfoListResponse) SetStatusCode(v int32) *GetAiRtcLicenseInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiRtcLicenseInfoListResponse) SetBody(v *GetAiRtcLicenseInfoListResponseBody) *GetAiRtcLicenseInfoListResponse {
	s.Body = v
	return s
}

func (s *GetAiRtcLicenseInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
