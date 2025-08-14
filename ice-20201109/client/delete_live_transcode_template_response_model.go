// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveTranscodeTemplateResponseBody) *DeleteLiveTranscodeTemplateResponse
	GetBody() *DeleteLiveTranscodeTemplateResponseBody
}

type DeleteLiveTranscodeTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveTranscodeTemplateResponse) GetBody() *DeleteLiveTranscodeTemplateResponseBody {
	return s.Body
}

func (s *DeleteLiveTranscodeTemplateResponse) SetHeaders(v map[string]*string) *DeleteLiveTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveTranscodeTemplateResponse) SetStatusCode(v int32) *DeleteLiveTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveTranscodeTemplateResponse) SetBody(v *DeleteLiveTranscodeTemplateResponseBody) *DeleteLiveTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveTranscodeTemplateResponse) Validate() error {
	return dara.Validate(s)
}
