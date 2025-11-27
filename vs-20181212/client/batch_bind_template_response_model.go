// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchBindTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchBindTemplateResponse
	GetStatusCode() *int32
	SetBody(v *BatchBindTemplateResponseBody) *BatchBindTemplateResponse
	GetBody() *BatchBindTemplateResponseBody
}

type BatchBindTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchBindTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchBindTemplateResponse) GetBody() *BatchBindTemplateResponseBody {
	return s.Body
}

func (s *BatchBindTemplateResponse) SetHeaders(v map[string]*string) *BatchBindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BatchBindTemplateResponse) SetStatusCode(v int32) *BatchBindTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindTemplateResponse) SetBody(v *BatchBindTemplateResponseBody) *BatchBindTemplateResponse {
	s.Body = v
	return s
}

func (s *BatchBindTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
