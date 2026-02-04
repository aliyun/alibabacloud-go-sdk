// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchEnisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGrantVSwitchEnisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGrantVSwitchEnisResponse
	GetStatusCode() *int32
	SetBody(v *ListGrantVSwitchEnisResponseBody) *ListGrantVSwitchEnisResponse
	GetBody() *ListGrantVSwitchEnisResponseBody
}

type ListGrantVSwitchEnisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGrantVSwitchEnisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGrantVSwitchEnisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchEnisResponse) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchEnisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGrantVSwitchEnisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGrantVSwitchEnisResponse) GetBody() *ListGrantVSwitchEnisResponseBody {
	return s.Body
}

func (s *ListGrantVSwitchEnisResponse) SetHeaders(v map[string]*string) *ListGrantVSwitchEnisResponse {
	s.Headers = v
	return s
}

func (s *ListGrantVSwitchEnisResponse) SetStatusCode(v int32) *ListGrantVSwitchEnisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGrantVSwitchEnisResponse) SetBody(v *ListGrantVSwitchEnisResponseBody) *ListGrantVSwitchEnisResponse {
	s.Body = v
	return s
}

func (s *ListGrantVSwitchEnisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
