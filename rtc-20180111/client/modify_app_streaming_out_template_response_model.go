// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppStreamingOutTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppStreamingOutTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppStreamingOutTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppStreamingOutTemplateResponseBody) *ModifyAppStreamingOutTemplateResponse
	GetBody() *ModifyAppStreamingOutTemplateResponseBody
}

type ModifyAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppStreamingOutTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppStreamingOutTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppStreamingOutTemplateResponse) GetBody() *ModifyAppStreamingOutTemplateResponseBody {
	return s.Body
}

func (s *ModifyAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *ModifyAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponse) SetStatusCode(v int32) *ModifyAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponse) SetBody(v *ModifyAppStreamingOutTemplateResponseBody) *ModifyAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponse) Validate() error {
	return dara.Validate(s)
}
