// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchTemplateResponseBody) *CreateBatchTemplateResponse
	GetBody() *CreateBatchTemplateResponseBody
}

type CreateBatchTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchTemplateResponse) GetBody() *CreateBatchTemplateResponseBody {
	return s.Body
}

func (s *CreateBatchTemplateResponse) SetHeaders(v map[string]*string) *CreateBatchTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchTemplateResponse) SetStatusCode(v int32) *CreateBatchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchTemplateResponse) SetBody(v *CreateBatchTemplateResponseBody) *CreateBatchTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateBatchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
