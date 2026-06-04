// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchParseDocumentLayoutTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchParseDocumentLayoutTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchParseDocumentLayoutTaskResponse
	GetStatusCode() *int32
	SetBody(v *FetchParseDocumentLayoutTaskResponseBody) *FetchParseDocumentLayoutTaskResponse
	GetBody() *FetchParseDocumentLayoutTaskResponseBody
}

type FetchParseDocumentLayoutTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchParseDocumentLayoutTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchParseDocumentLayoutTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchParseDocumentLayoutTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchParseDocumentLayoutTaskResponse) GetBody() *FetchParseDocumentLayoutTaskResponseBody {
	return s.Body
}

func (s *FetchParseDocumentLayoutTaskResponse) SetHeaders(v map[string]*string) *FetchParseDocumentLayoutTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponse) SetStatusCode(v int32) *FetchParseDocumentLayoutTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponse) SetBody(v *FetchParseDocumentLayoutTaskResponseBody) *FetchParseDocumentLayoutTaskResponse {
	s.Body = v
	return s
}

func (s *FetchParseDocumentLayoutTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
