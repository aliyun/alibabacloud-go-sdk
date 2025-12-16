// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompileSortScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompileSortScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompileSortScriptResponse
	GetStatusCode() *int32
	SetBody(v *CompileSortScriptResponseBody) *CompileSortScriptResponse
	GetBody() *CompileSortScriptResponseBody
}

type CompileSortScriptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompileSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompileSortScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s CompileSortScriptResponse) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompileSortScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompileSortScriptResponse) GetBody() *CompileSortScriptResponseBody {
	return s.Body
}

func (s *CompileSortScriptResponse) SetHeaders(v map[string]*string) *CompileSortScriptResponse {
	s.Headers = v
	return s
}

func (s *CompileSortScriptResponse) SetStatusCode(v int32) *CompileSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *CompileSortScriptResponse) SetBody(v *CompileSortScriptResponseBody) *CompileSortScriptResponse {
	s.Body = v
	return s
}

func (s *CompileSortScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
