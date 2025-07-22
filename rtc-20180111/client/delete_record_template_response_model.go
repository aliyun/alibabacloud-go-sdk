// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecordTemplateResponseBody) *DeleteRecordTemplateResponse
	GetBody() *DeleteRecordTemplateResponseBody
}

type DeleteRecordTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecordTemplateResponse) GetBody() *DeleteRecordTemplateResponseBody {
	return s.Body
}

func (s *DeleteRecordTemplateResponse) SetHeaders(v map[string]*string) *DeleteRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordTemplateResponse) SetStatusCode(v int32) *DeleteRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordTemplateResponse) SetBody(v *DeleteRecordTemplateResponseBody) *DeleteRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteRecordTemplateResponse) Validate() error {
	return dara.Validate(s)
}
