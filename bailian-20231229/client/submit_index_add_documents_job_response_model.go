// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexAddDocumentsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIndexAddDocumentsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIndexAddDocumentsJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIndexAddDocumentsJobResponseBody) *SubmitIndexAddDocumentsJobResponse
	GetBody() *SubmitIndexAddDocumentsJobResponseBody
}

type SubmitIndexAddDocumentsJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIndexAddDocumentsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIndexAddDocumentsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIndexAddDocumentsJobResponse) GetBody() *SubmitIndexAddDocumentsJobResponseBody {
	return s.Body
}

func (s *SubmitIndexAddDocumentsJobResponse) SetHeaders(v map[string]*string) *SubmitIndexAddDocumentsJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponse) SetStatusCode(v int32) *SubmitIndexAddDocumentsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponse) SetBody(v *SubmitIndexAddDocumentsJobResponseBody) *SubmitIndexAddDocumentsJobResponse {
	s.Body = v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
