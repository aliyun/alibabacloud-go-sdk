// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishScriptResponse
	GetStatusCode() *int32
	SetBody(v *PublishScriptResponseBody) *PublishScriptResponse
	GetBody() *PublishScriptResponseBody
}

type PublishScriptResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptResponse) GoString() string {
	return s.String()
}

func (s *PublishScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishScriptResponse) GetBody() *PublishScriptResponseBody {
	return s.Body
}

func (s *PublishScriptResponse) SetHeaders(v map[string]*string) *PublishScriptResponse {
	s.Headers = v
	return s
}

func (s *PublishScriptResponse) SetStatusCode(v int32) *PublishScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishScriptResponse) SetBody(v *PublishScriptResponseBody) *PublishScriptResponse {
	s.Body = v
	return s
}

func (s *PublishScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
