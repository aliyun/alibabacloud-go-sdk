// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse
	GetBody() *ModifyRemarkResponseBody
}

type ModifyRemarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRemarkResponse) GetBody() *ModifyRemarkResponseBody {
	return s.Body
}

func (s *ModifyRemarkResponse) SetHeaders(v map[string]*string) *ModifyRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyRemarkResponse) SetStatusCode(v int32) *ModifyRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRemarkResponse) SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse {
	s.Body = v
	return s
}

func (s *ModifyRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
