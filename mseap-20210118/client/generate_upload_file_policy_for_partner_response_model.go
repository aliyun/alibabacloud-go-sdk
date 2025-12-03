// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUploadFilePolicyForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUploadFilePolicyForPartnerResponseBody) *GenerateUploadFilePolicyForPartnerResponse
	GetBody() *GenerateUploadFilePolicyForPartnerResponseBody
}

type GenerateUploadFilePolicyForPartnerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadFilePolicyForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUploadFilePolicyForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUploadFilePolicyForPartnerResponse) GetBody() *GenerateUploadFilePolicyForPartnerResponseBody {
	return s.Body
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetBody(v *GenerateUploadFilePolicyForPartnerResponseBody) *GenerateUploadFilePolicyForPartnerResponse {
	s.Body = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
