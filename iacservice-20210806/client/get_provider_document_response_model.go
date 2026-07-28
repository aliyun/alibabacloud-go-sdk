// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProviderDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProviderDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProviderDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetProviderDocumentResponseBody) *GetProviderDocumentResponse
	GetBody() *GetProviderDocumentResponseBody
}

type GetProviderDocumentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProviderDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProviderDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProviderDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetProviderDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProviderDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProviderDocumentResponse) GetBody() *GetProviderDocumentResponseBody {
	return s.Body
}

func (s *GetProviderDocumentResponse) SetHeaders(v map[string]*string) *GetProviderDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetProviderDocumentResponse) SetStatusCode(v int32) *GetProviderDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProviderDocumentResponse) SetBody(v *GetProviderDocumentResponseBody) *GetProviderDocumentResponse {
	s.Body = v
	return s
}

func (s *GetProviderDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
