// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentThemeUploadSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentThemeUploadSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentThemeUploadSignatureResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentThemeUploadSignatureResponseBody) *GetDataAgentThemeUploadSignatureResponse
	GetBody() *GetDataAgentThemeUploadSignatureResponseBody
}

type GetDataAgentThemeUploadSignatureResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentThemeUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentThemeUploadSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentThemeUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentThemeUploadSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentThemeUploadSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentThemeUploadSignatureResponse) GetBody() *GetDataAgentThemeUploadSignatureResponseBody {
	return s.Body
}

func (s *GetDataAgentThemeUploadSignatureResponse) SetHeaders(v map[string]*string) *GetDataAgentThemeUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponse) SetStatusCode(v int32) *GetDataAgentThemeUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponse) SetBody(v *GetDataAgentThemeUploadSignatureResponseBody) *GetDataAgentThemeUploadSignatureResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
