// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceUploadSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourceUploadSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourceUploadSignatureResponse
	GetStatusCode() *int32
	SetBody(v *GetSourceUploadSignatureResponseBody) *GetSourceUploadSignatureResponse
	GetBody() *GetSourceUploadSignatureResponseBody
}

type GetSourceUploadSignatureResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourceUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourceUploadSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourceUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetSourceUploadSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourceUploadSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourceUploadSignatureResponse) GetBody() *GetSourceUploadSignatureResponseBody {
	return s.Body
}

func (s *GetSourceUploadSignatureResponse) SetHeaders(v map[string]*string) *GetSourceUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetSourceUploadSignatureResponse) SetStatusCode(v int32) *GetSourceUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceUploadSignatureResponse) SetBody(v *GetSourceUploadSignatureResponseBody) *GetSourceUploadSignatureResponse {
	s.Body = v
	return s
}

func (s *GetSourceUploadSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
