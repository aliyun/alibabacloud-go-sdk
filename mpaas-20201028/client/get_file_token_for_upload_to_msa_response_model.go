// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTokenForUploadToMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileTokenForUploadToMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileTokenForUploadToMsaResponse
	GetStatusCode() *int32
	SetBody(v *GetFileTokenForUploadToMsaResponseBody) *GetFileTokenForUploadToMsaResponse
	GetBody() *GetFileTokenForUploadToMsaResponseBody
}

type GetFileTokenForUploadToMsaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileTokenForUploadToMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileTokenForUploadToMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileTokenForUploadToMsaResponse) GoString() string {
	return s.String()
}

func (s *GetFileTokenForUploadToMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileTokenForUploadToMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileTokenForUploadToMsaResponse) GetBody() *GetFileTokenForUploadToMsaResponseBody {
	return s.Body
}

func (s *GetFileTokenForUploadToMsaResponse) SetHeaders(v map[string]*string) *GetFileTokenForUploadToMsaResponse {
	s.Headers = v
	return s
}

func (s *GetFileTokenForUploadToMsaResponse) SetStatusCode(v int32) *GetFileTokenForUploadToMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponse) SetBody(v *GetFileTokenForUploadToMsaResponseBody) *GetFileTokenForUploadToMsaResponse {
	s.Body = v
	return s
}

func (s *GetFileTokenForUploadToMsaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
