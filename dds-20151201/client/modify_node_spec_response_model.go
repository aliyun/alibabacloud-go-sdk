// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodeSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodeSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodeSpecResponseBody) *ModifyNodeSpecResponse
	GetBody() *ModifyNodeSpecResponseBody
}

type ModifyNodeSpecResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodeSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodeSpecResponse) GetBody() *ModifyNodeSpecResponseBody {
	return s.Body
}

func (s *ModifyNodeSpecResponse) SetHeaders(v map[string]*string) *ModifyNodeSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeSpecResponse) SetStatusCode(v int32) *ModifyNodeSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeSpecResponse) SetBody(v *ModifyNodeSpecResponseBody) *ModifyNodeSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyNodeSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
