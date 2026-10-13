// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomScriptTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomScriptTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomScriptTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomScriptTemplatesResponseBody) *ListCustomScriptTemplatesResponse
	GetBody() *ListCustomScriptTemplatesResponseBody
}

type ListCustomScriptTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomScriptTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomScriptTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomScriptTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomScriptTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomScriptTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomScriptTemplatesResponse) GetBody() *ListCustomScriptTemplatesResponseBody {
	return s.Body
}

func (s *ListCustomScriptTemplatesResponse) SetHeaders(v map[string]*string) *ListCustomScriptTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomScriptTemplatesResponse) SetStatusCode(v int32) *ListCustomScriptTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomScriptTemplatesResponse) SetBody(v *ListCustomScriptTemplatesResponseBody) *ListCustomScriptTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListCustomScriptTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
