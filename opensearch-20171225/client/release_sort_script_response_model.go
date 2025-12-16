// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSortScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseSortScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseSortScriptResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseSortScriptResponseBody) *ReleaseSortScriptResponse
	GetBody() *ReleaseSortScriptResponseBody
}

type ReleaseSortScriptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSortScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSortScriptResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseSortScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseSortScriptResponse) GetBody() *ReleaseSortScriptResponseBody {
	return s.Body
}

func (s *ReleaseSortScriptResponse) SetHeaders(v map[string]*string) *ReleaseSortScriptResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSortScriptResponse) SetStatusCode(v int32) *ReleaseSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSortScriptResponse) SetBody(v *ReleaseSortScriptResponseBody) *ReleaseSortScriptResponse {
	s.Body = v
	return s
}

func (s *ReleaseSortScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
