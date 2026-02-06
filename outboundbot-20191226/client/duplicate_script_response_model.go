// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DuplicateScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DuplicateScriptResponse
	GetStatusCode() *int32
	SetBody(v *DuplicateScriptResponseBody) *DuplicateScriptResponse
	GetBody() *DuplicateScriptResponseBody
}

type DuplicateScriptResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DuplicateScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DuplicateScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s DuplicateScriptResponse) GoString() string {
	return s.String()
}

func (s *DuplicateScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DuplicateScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DuplicateScriptResponse) GetBody() *DuplicateScriptResponseBody {
	return s.Body
}

func (s *DuplicateScriptResponse) SetHeaders(v map[string]*string) *DuplicateScriptResponse {
	s.Headers = v
	return s
}

func (s *DuplicateScriptResponse) SetStatusCode(v int32) *DuplicateScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DuplicateScriptResponse) SetBody(v *DuplicateScriptResponseBody) *DuplicateScriptResponse {
	s.Body = v
	return s
}

func (s *DuplicateScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
