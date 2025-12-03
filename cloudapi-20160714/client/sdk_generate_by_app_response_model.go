// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdkGenerateByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdkGenerateByAppResponse
	GetStatusCode() *int32
	SetBody(v *SdkGenerateByAppResponseBody) *SdkGenerateByAppResponse
	GetBody() *SdkGenerateByAppResponseBody
}

type SdkGenerateByAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdkGenerateByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdkGenerateByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByAppResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdkGenerateByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdkGenerateByAppResponse) GetBody() *SdkGenerateByAppResponseBody {
	return s.Body
}

func (s *SdkGenerateByAppResponse) SetHeaders(v map[string]*string) *SdkGenerateByAppResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateByAppResponse) SetStatusCode(v int32) *SdkGenerateByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateByAppResponse) SetBody(v *SdkGenerateByAppResponseBody) *SdkGenerateByAppResponse {
	s.Body = v
	return s
}

func (s *SdkGenerateByAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
