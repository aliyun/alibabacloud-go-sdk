// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppSiteValidationFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAppSiteValidationFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAppSiteValidationFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadAppSiteValidationFileResponseBody) *UploadAppSiteValidationFileResponse
	GetBody() *UploadAppSiteValidationFileResponseBody
}

type UploadAppSiteValidationFileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAppSiteValidationFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAppSiteValidationFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAppSiteValidationFileResponse) GoString() string {
	return s.String()
}

func (s *UploadAppSiteValidationFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAppSiteValidationFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAppSiteValidationFileResponse) GetBody() *UploadAppSiteValidationFileResponseBody {
	return s.Body
}

func (s *UploadAppSiteValidationFileResponse) SetHeaders(v map[string]*string) *UploadAppSiteValidationFileResponse {
	s.Headers = v
	return s
}

func (s *UploadAppSiteValidationFileResponse) SetStatusCode(v int32) *UploadAppSiteValidationFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAppSiteValidationFileResponse) SetBody(v *UploadAppSiteValidationFileResponseBody) *UploadAppSiteValidationFileResponse {
	s.Body = v
	return s
}

func (s *UploadAppSiteValidationFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
