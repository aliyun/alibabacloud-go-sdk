// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBatchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBatchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBatchTemplateResponseBody) *UpdateBatchTemplateResponse
	GetBody() *UpdateBatchTemplateResponseBody
}

type UpdateBatchTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBatchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBatchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBatchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBatchTemplateResponse) GetBody() *UpdateBatchTemplateResponseBody {
	return s.Body
}

func (s *UpdateBatchTemplateResponse) SetHeaders(v map[string]*string) *UpdateBatchTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchTemplateResponse) SetStatusCode(v int32) *UpdateBatchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchTemplateResponse) SetBody(v *UpdateBatchTemplateResponseBody) *UpdateBatchTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateBatchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
