// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentCompareResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentCompareResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentCompareResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentCompareResultResponseBody) *GetDocumentCompareResultResponse
	GetBody() *GetDocumentCompareResultResponseBody
}

type GetDocumentCompareResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentCompareResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentCompareResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentCompareResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentCompareResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentCompareResultResponse) GetBody() *GetDocumentCompareResultResponseBody {
	return s.Body
}

func (s *GetDocumentCompareResultResponse) SetHeaders(v map[string]*string) *GetDocumentCompareResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentCompareResultResponse) SetStatusCode(v int32) *GetDocumentCompareResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentCompareResultResponse) SetBody(v *GetDocumentCompareResultResponseBody) *GetDocumentCompareResultResponse {
	s.Body = v
	return s
}

func (s *GetDocumentCompareResultResponse) Validate() error {
	return dara.Validate(s)
}
