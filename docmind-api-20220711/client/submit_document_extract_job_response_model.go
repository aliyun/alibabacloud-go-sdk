// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocumentExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocumentExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocumentExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocumentExtractJobResponseBody) *SubmitDocumentExtractJobResponse
	GetBody() *SubmitDocumentExtractJobResponseBody
}

type SubmitDocumentExtractJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocumentExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocumentExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocumentExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocumentExtractJobResponse) GetBody() *SubmitDocumentExtractJobResponseBody {
	return s.Body
}

func (s *SubmitDocumentExtractJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentExtractJobResponse) SetStatusCode(v int32) *SubmitDocumentExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentExtractJobResponse) SetBody(v *SubmitDocumentExtractJobResponseBody) *SubmitDocumentExtractJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDocumentExtractJobResponse) Validate() error {
	return dara.Validate(s)
}
