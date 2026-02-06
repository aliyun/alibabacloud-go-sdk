// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportScriptResponse
	GetStatusCode() *int32
	SetBody(v *ImportScriptResponseBody) *ImportScriptResponse
	GetBody() *ImportScriptResponseBody
}

type ImportScriptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportScriptResponse) GoString() string {
	return s.String()
}

func (s *ImportScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportScriptResponse) GetBody() *ImportScriptResponseBody {
	return s.Body
}

func (s *ImportScriptResponse) SetHeaders(v map[string]*string) *ImportScriptResponse {
	s.Headers = v
	return s
}

func (s *ImportScriptResponse) SetStatusCode(v int32) *ImportScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportScriptResponse) SetBody(v *ImportScriptResponseBody) *ImportScriptResponse {
	s.Body = v
	return s
}

func (s *ImportScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
