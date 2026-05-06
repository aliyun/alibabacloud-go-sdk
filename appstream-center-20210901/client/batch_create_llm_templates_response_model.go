// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateLlmTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateLlmTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateLlmTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateLlmTemplatesResponseBody) *BatchCreateLlmTemplatesResponse
	GetBody() *BatchCreateLlmTemplatesResponseBody
}

type BatchCreateLlmTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateLlmTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateLlmTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateLlmTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateLlmTemplatesResponse) GetBody() *BatchCreateLlmTemplatesResponseBody {
	return s.Body
}

func (s *BatchCreateLlmTemplatesResponse) SetHeaders(v map[string]*string) *BatchCreateLlmTemplatesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateLlmTemplatesResponse) SetStatusCode(v int32) *BatchCreateLlmTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponse) SetBody(v *BatchCreateLlmTemplatesResponseBody) *BatchCreateLlmTemplatesResponse {
	s.Body = v
	return s
}

func (s *BatchCreateLlmTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
