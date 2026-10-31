// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBatchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBatchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBatchTemplateResponseBody) *SubmitBatchTemplateResponse
	GetBody() *SubmitBatchTemplateResponseBody
}

type SubmitBatchTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBatchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBatchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateResponse) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBatchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBatchTemplateResponse) GetBody() *SubmitBatchTemplateResponseBody {
	return s.Body
}

func (s *SubmitBatchTemplateResponse) SetHeaders(v map[string]*string) *SubmitBatchTemplateResponse {
	s.Headers = v
	return s
}

func (s *SubmitBatchTemplateResponse) SetStatusCode(v int32) *SubmitBatchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBatchTemplateResponse) SetBody(v *SubmitBatchTemplateResponseBody) *SubmitBatchTemplateResponse {
	s.Body = v
	return s
}

func (s *SubmitBatchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
