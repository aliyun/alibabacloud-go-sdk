// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetouchSkinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetouchSkinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetouchSkinResponse
	GetStatusCode() *int32
	SetBody(v *RetouchSkinResponseBody) *RetouchSkinResponse
	GetBody() *RetouchSkinResponseBody
}

type RetouchSkinResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetouchSkinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetouchSkinResponse) String() string {
	return dara.Prettify(s)
}

func (s RetouchSkinResponse) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetouchSkinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetouchSkinResponse) GetBody() *RetouchSkinResponseBody {
	return s.Body
}

func (s *RetouchSkinResponse) SetHeaders(v map[string]*string) *RetouchSkinResponse {
	s.Headers = v
	return s
}

func (s *RetouchSkinResponse) SetStatusCode(v int32) *RetouchSkinResponse {
	s.StatusCode = &v
	return s
}

func (s *RetouchSkinResponse) SetBody(v *RetouchSkinResponseBody) *RetouchSkinResponse {
	s.Body = v
	return s
}

func (s *RetouchSkinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
