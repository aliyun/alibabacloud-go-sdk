// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodeNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodeNumberResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodeNumberResponseBody) *ModifyNodeNumberResponse
	GetBody() *ModifyNodeNumberResponseBody
}

type ModifyNodeNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodeNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodeNumberResponse) GetBody() *ModifyNodeNumberResponseBody {
	return s.Body
}

func (s *ModifyNodeNumberResponse) SetHeaders(v map[string]*string) *ModifyNodeNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeNumberResponse) SetStatusCode(v int32) *ModifyNodeNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeNumberResponse) SetBody(v *ModifyNodeNumberResponseBody) *ModifyNodeNumberResponse {
	s.Body = v
	return s
}

func (s *ModifyNodeNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
