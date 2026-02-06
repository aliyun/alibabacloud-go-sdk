// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptForDebugResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishScriptForDebugResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishScriptForDebugResponse
	GetStatusCode() *int32
	SetBody(v *PublishScriptForDebugResponseBody) *PublishScriptForDebugResponse
	GetBody() *PublishScriptForDebugResponseBody
}

type PublishScriptForDebugResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishScriptForDebugResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishScriptForDebugResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptForDebugResponse) GoString() string {
	return s.String()
}

func (s *PublishScriptForDebugResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishScriptForDebugResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishScriptForDebugResponse) GetBody() *PublishScriptForDebugResponseBody {
	return s.Body
}

func (s *PublishScriptForDebugResponse) SetHeaders(v map[string]*string) *PublishScriptForDebugResponse {
	s.Headers = v
	return s
}

func (s *PublishScriptForDebugResponse) SetStatusCode(v int32) *PublishScriptForDebugResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishScriptForDebugResponse) SetBody(v *PublishScriptForDebugResponseBody) *PublishScriptForDebugResponse {
	s.Body = v
	return s
}

func (s *PublishScriptForDebugResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
