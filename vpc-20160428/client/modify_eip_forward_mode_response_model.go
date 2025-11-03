// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipForwardModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEipForwardModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEipForwardModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEipForwardModeResponseBody) *ModifyEipForwardModeResponse
	GetBody() *ModifyEipForwardModeResponseBody
}

type ModifyEipForwardModeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEipForwardModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEipForwardModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipForwardModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyEipForwardModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEipForwardModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEipForwardModeResponse) GetBody() *ModifyEipForwardModeResponseBody {
	return s.Body
}

func (s *ModifyEipForwardModeResponse) SetHeaders(v map[string]*string) *ModifyEipForwardModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyEipForwardModeResponse) SetStatusCode(v int32) *ModifyEipForwardModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEipForwardModeResponse) SetBody(v *ModifyEipForwardModeResponseBody) *ModifyEipForwardModeResponse {
	s.Body = v
	return s
}

func (s *ModifyEipForwardModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
