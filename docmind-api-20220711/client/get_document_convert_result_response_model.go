// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentConvertResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentConvertResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentConvertResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentConvertResultResponseBody) *GetDocumentConvertResultResponse
	GetBody() *GetDocumentConvertResultResponseBody
}

type GetDocumentConvertResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentConvertResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentConvertResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentConvertResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentConvertResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentConvertResultResponse) GetBody() *GetDocumentConvertResultResponseBody {
	return s.Body
}

func (s *GetDocumentConvertResultResponse) SetHeaders(v map[string]*string) *GetDocumentConvertResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentConvertResultResponse) SetStatusCode(v int32) *GetDocumentConvertResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentConvertResultResponse) SetBody(v *GetDocumentConvertResultResponseBody) *GetDocumentConvertResultResponse {
	s.Body = v
	return s
}

func (s *GetDocumentConvertResultResponse) Validate() error {
	return dara.Validate(s)
}
