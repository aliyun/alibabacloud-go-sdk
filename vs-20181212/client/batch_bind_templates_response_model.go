// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchBindTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchBindTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *BatchBindTemplatesResponseBody) *BatchBindTemplatesResponse
	GetBody() *BatchBindTemplatesResponseBody
}

type BatchBindTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplatesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchBindTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchBindTemplatesResponse) GetBody() *BatchBindTemplatesResponseBody {
	return s.Body
}

func (s *BatchBindTemplatesResponse) SetHeaders(v map[string]*string) *BatchBindTemplatesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindTemplatesResponse) SetStatusCode(v int32) *BatchBindTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindTemplatesResponse) SetBody(v *BatchBindTemplatesResponseBody) *BatchBindTemplatesResponse {
	s.Body = v
	return s
}

func (s *BatchBindTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
