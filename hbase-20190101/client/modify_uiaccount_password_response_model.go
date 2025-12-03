// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUIAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUIAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUIAccountPasswordResponseBody) *ModifyUIAccountPasswordResponse
	GetBody() *ModifyUIAccountPasswordResponseBody
}

type ModifyUIAccountPasswordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUIAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUIAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUIAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUIAccountPasswordResponse) GetBody() *ModifyUIAccountPasswordResponseBody {
	return s.Body
}

func (s *ModifyUIAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyUIAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUIAccountPasswordResponse) SetStatusCode(v int32) *ModifyUIAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUIAccountPasswordResponse) SetBody(v *ModifyUIAccountPasswordResponseBody) *ModifyUIAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyUIAccountPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
