// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaDocumentResponseBody) *GetLumaDocumentResponse
	GetBody() *GetLumaDocumentResponseBody
}

type GetLumaDocumentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaDocumentResponse) GetBody() *GetLumaDocumentResponseBody {
	return s.Body
}

func (s *GetLumaDocumentResponse) SetHeaders(v map[string]*string) *GetLumaDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetLumaDocumentResponse) SetStatusCode(v int32) *GetLumaDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaDocumentResponse) SetBody(v *GetLumaDocumentResponseBody) *GetLumaDocumentResponse {
	s.Body = v
	return s
}

func (s *GetLumaDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
