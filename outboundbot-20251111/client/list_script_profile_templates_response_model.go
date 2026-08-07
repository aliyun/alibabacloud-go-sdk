// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptProfileTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptProfileTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptProfileTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptProfileTemplatesResponseBody) *ListScriptProfileTemplatesResponse
	GetBody() *ListScriptProfileTemplatesResponseBody
}

type ListScriptProfileTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptProfileTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptProfileTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptProfileTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptProfileTemplatesResponse) GetBody() *ListScriptProfileTemplatesResponseBody {
	return s.Body
}

func (s *ListScriptProfileTemplatesResponse) SetHeaders(v map[string]*string) *ListScriptProfileTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListScriptProfileTemplatesResponse) SetStatusCode(v int32) *ListScriptProfileTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptProfileTemplatesResponse) SetBody(v *ListScriptProfileTemplatesResponseBody) *ListScriptProfileTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListScriptProfileTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
