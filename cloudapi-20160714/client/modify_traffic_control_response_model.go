// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTrafficControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTrafficControlResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTrafficControlResponseBody) *ModifyTrafficControlResponse
	GetBody() *ModifyTrafficControlResponseBody
}

type ModifyTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrafficControlResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTrafficControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTrafficControlResponse) GetBody() *ModifyTrafficControlResponseBody {
	return s.Body
}

func (s *ModifyTrafficControlResponse) SetHeaders(v map[string]*string) *ModifyTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrafficControlResponse) SetStatusCode(v int32) *ModifyTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrafficControlResponse) SetBody(v *ModifyTrafficControlResponseBody) *ModifyTrafficControlResponse {
	s.Body = v
	return s
}

func (s *ModifyTrafficControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
