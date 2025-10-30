// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiDocumentResponseBody) *GetDataServiceApiDocumentResponse
	GetBody() *GetDataServiceApiDocumentResponseBody
}

type GetDataServiceApiDocumentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiDocumentResponse) GetBody() *GetDataServiceApiDocumentResponseBody {
	return s.Body
}

func (s *GetDataServiceApiDocumentResponse) SetHeaders(v map[string]*string) *GetDataServiceApiDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiDocumentResponse) SetStatusCode(v int32) *GetDataServiceApiDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiDocumentResponse) SetBody(v *GetDataServiceApiDocumentResponseBody) *GetDataServiceApiDocumentResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
