// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSmartCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocSmartCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocSmartCardResponse
	GetStatusCode() *int32
	SetBody(v *RunDocSmartCardResponseBody) *RunDocSmartCardResponse
	GetBody() *RunDocSmartCardResponseBody
}

type RunDocSmartCardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDocSmartCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocSmartCardResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponse) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocSmartCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocSmartCardResponse) GetBody() *RunDocSmartCardResponseBody {
	return s.Body
}

func (s *RunDocSmartCardResponse) SetHeaders(v map[string]*string) *RunDocSmartCardResponse {
	s.Headers = v
	return s
}

func (s *RunDocSmartCardResponse) SetStatusCode(v int32) *RunDocSmartCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocSmartCardResponse) SetBody(v *RunDocSmartCardResponseBody) *RunDocSmartCardResponse {
	s.Body = v
	return s
}

func (s *RunDocSmartCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
