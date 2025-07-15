// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByAppForRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdkGenerateByAppForRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdkGenerateByAppForRegionResponse
	GetStatusCode() *int32
	SetBody(v *SdkGenerateByAppForRegionResponseBody) *SdkGenerateByAppForRegionResponse
	GetBody() *SdkGenerateByAppForRegionResponseBody
}

type SdkGenerateByAppForRegionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdkGenerateByAppForRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdkGenerateByAppForRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByAppForRegionResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppForRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdkGenerateByAppForRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdkGenerateByAppForRegionResponse) GetBody() *SdkGenerateByAppForRegionResponseBody {
	return s.Body
}

func (s *SdkGenerateByAppForRegionResponse) SetHeaders(v map[string]*string) *SdkGenerateByAppForRegionResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateByAppForRegionResponse) SetStatusCode(v int32) *SdkGenerateByAppForRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateByAppForRegionResponse) SetBody(v *SdkGenerateByAppForRegionResponseBody) *SdkGenerateByAppForRegionResponse {
	s.Body = v
	return s
}

func (s *SdkGenerateByAppForRegionResponse) Validate() error {
	return dara.Validate(s)
}
