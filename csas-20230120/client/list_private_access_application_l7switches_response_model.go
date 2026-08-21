// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationL7SwitchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessApplicationL7SwitchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessApplicationL7SwitchesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessApplicationL7SwitchesResponseBody) *ListPrivateAccessApplicationL7SwitchesResponse
	GetBody() *ListPrivateAccessApplicationL7SwitchesResponseBody
}

type ListPrivateAccessApplicationL7SwitchesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessApplicationL7SwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessApplicationL7SwitchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationL7SwitchesResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) GetBody() *ListPrivateAccessApplicationL7SwitchesResponseBody {
	return s.Body
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) SetHeaders(v map[string]*string) *ListPrivateAccessApplicationL7SwitchesResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) SetStatusCode(v int32) *ListPrivateAccessApplicationL7SwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) SetBody(v *ListPrivateAccessApplicationL7SwitchesResponseBody) *ListPrivateAccessApplicationL7SwitchesResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
