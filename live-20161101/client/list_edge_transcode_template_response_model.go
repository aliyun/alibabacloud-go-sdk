// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeTranscodeTemplateResponseBody) *ListEdgeTranscodeTemplateResponse
	GetBody() *ListEdgeTranscodeTemplateResponseBody
}

type ListEdgeTranscodeTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeTranscodeTemplateResponse) GetBody() *ListEdgeTranscodeTemplateResponseBody {
	return s.Body
}

func (s *ListEdgeTranscodeTemplateResponse) SetHeaders(v map[string]*string) *ListEdgeTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeTranscodeTemplateResponse) SetStatusCode(v int32) *ListEdgeTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponse) SetBody(v *ListEdgeTranscodeTemplateResponseBody) *ListEdgeTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *ListEdgeTranscodeTemplateResponse) Validate() error {
	return dara.Validate(s)
}
