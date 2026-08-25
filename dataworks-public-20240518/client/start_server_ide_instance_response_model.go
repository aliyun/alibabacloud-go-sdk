// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartServerIdeInstanceResponseBody) *StartServerIdeInstanceResponse
	GetBody() *StartServerIdeInstanceResponseBody
}

type StartServerIdeInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartServerIdeInstanceResponse) GetBody() *StartServerIdeInstanceResponseBody {
	return s.Body
}

func (s *StartServerIdeInstanceResponse) SetHeaders(v map[string]*string) *StartServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServerIdeInstanceResponse) SetStatusCode(v int32) *StartServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServerIdeInstanceResponse) SetBody(v *StartServerIdeInstanceResponseBody) *StartServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *StartServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
