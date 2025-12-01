// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentSummaryResponseBody) *GetDocumentSummaryResponse
	GetBody() *GetDocumentSummaryResponseBody
}

type GetDocumentSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentSummaryResponse) GetBody() *GetDocumentSummaryResponseBody {
	return s.Body
}

func (s *GetDocumentSummaryResponse) SetHeaders(v map[string]*string) *GetDocumentSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentSummaryResponse) SetStatusCode(v int32) *GetDocumentSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponse) SetBody(v *GetDocumentSummaryResponseBody) *GetDocumentSummaryResponse {
	s.Body = v
	return s
}

func (s *GetDocumentSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
