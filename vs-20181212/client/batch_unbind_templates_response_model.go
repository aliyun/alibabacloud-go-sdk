// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUnbindTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUnbindTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUnbindTemplatesResponseBody) *BatchUnbindTemplatesResponse
	GetBody() *BatchUnbindTemplatesResponseBody
}

type BatchUnbindTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUnbindTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUnbindTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplatesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUnbindTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUnbindTemplatesResponse) GetBody() *BatchUnbindTemplatesResponseBody {
	return s.Body
}

func (s *BatchUnbindTemplatesResponse) SetHeaders(v map[string]*string) *BatchUnbindTemplatesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindTemplatesResponse) SetStatusCode(v int32) *BatchUnbindTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUnbindTemplatesResponse) SetBody(v *BatchUnbindTemplatesResponseBody) *BatchUnbindTemplatesResponse {
	s.Body = v
	return s
}

func (s *BatchUnbindTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
