// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVncPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVncPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVncPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVncPasswordResponseBody) *ModifyInstanceVncPasswordResponse
	GetBody() *ModifyInstanceVncPasswordResponseBody
}

type ModifyInstanceVncPasswordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVncPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVncPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVncPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVncPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVncPasswordResponse) GetBody() *ModifyInstanceVncPasswordResponseBody {
	return s.Body
}

func (s *ModifyInstanceVncPasswordResponse) SetHeaders(v map[string]*string) *ModifyInstanceVncPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVncPasswordResponse) SetStatusCode(v int32) *ModifyInstanceVncPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVncPasswordResponse) SetBody(v *ModifyInstanceVncPasswordResponseBody) *ModifyInstanceVncPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVncPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
