// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumberDistrictInfoTemplateDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNumberDistrictInfoTemplateDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNumberDistrictInfoTemplateDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) *GetNumberDistrictInfoTemplateDownloadUrlResponse
	GetBody() *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
}

type GetNumberDistrictInfoTemplateDownloadUrlResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNumberDistrictInfoTemplateDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNumberDistrictInfoTemplateDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNumberDistrictInfoTemplateDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) GetBody() *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	return s.Body
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) SetHeaders(v map[string]*string) *GetNumberDistrictInfoTemplateDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) SetStatusCode(v int32) *GetNumberDistrictInfoTemplateDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) SetBody(v *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) *GetNumberDistrictInfoTemplateDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
