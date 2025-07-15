// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUploadParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentUploadParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentUploadParametersResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentUploadParametersResponseBody) *GetDocumentUploadParametersResponse
	GetBody() *GetDocumentUploadParametersResponseBody
}

type GetDocumentUploadParametersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentUploadParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentUploadParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUploadParametersResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentUploadParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentUploadParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentUploadParametersResponse) GetBody() *GetDocumentUploadParametersResponseBody {
	return s.Body
}

func (s *GetDocumentUploadParametersResponse) SetHeaders(v map[string]*string) *GetDocumentUploadParametersResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentUploadParametersResponse) SetStatusCode(v int32) *GetDocumentUploadParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentUploadParametersResponse) SetBody(v *GetDocumentUploadParametersResponseBody) *GetDocumentUploadParametersResponse {
	s.Body = v
	return s
}

func (s *GetDocumentUploadParametersResponse) Validate() error {
	return dara.Validate(s)
}
