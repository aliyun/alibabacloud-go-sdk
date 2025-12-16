// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptFileNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScriptFileNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScriptFileNamesResponse
	GetStatusCode() *int32
	SetBody(v *GetScriptFileNamesResponseBody) *GetScriptFileNamesResponse
	GetBody() *GetScriptFileNamesResponseBody
}

type GetScriptFileNamesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptFileNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScriptFileNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScriptFileNamesResponse) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScriptFileNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScriptFileNamesResponse) GetBody() *GetScriptFileNamesResponseBody {
	return s.Body
}

func (s *GetScriptFileNamesResponse) SetHeaders(v map[string]*string) *GetScriptFileNamesResponse {
	s.Headers = v
	return s
}

func (s *GetScriptFileNamesResponse) SetStatusCode(v int32) *GetScriptFileNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptFileNamesResponse) SetBody(v *GetScriptFileNamesResponseBody) *GetScriptFileNamesResponse {
	s.Body = v
	return s
}

func (s *GetScriptFileNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
