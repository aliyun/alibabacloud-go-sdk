// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindEsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindEsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UnbindEsInstanceResponseBody) *UnbindEsInstanceResponse
	GetBody() *UnbindEsInstanceResponseBody
}

type UnbindEsInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindEsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindEsInstanceResponse) GetBody() *UnbindEsInstanceResponseBody {
	return s.Body
}

func (s *UnbindEsInstanceResponse) SetHeaders(v map[string]*string) *UnbindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnbindEsInstanceResponse) SetStatusCode(v int32) *UnbindEsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEsInstanceResponse) SetBody(v *UnbindEsInstanceResponseBody) *UnbindEsInstanceResponse {
	s.Body = v
	return s
}

func (s *UnbindEsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
