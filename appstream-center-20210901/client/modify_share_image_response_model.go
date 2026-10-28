// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShareImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyShareImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyShareImageResponse
	GetStatusCode() *int32
	SetBody(v *ModifyShareImageResponseBody) *ModifyShareImageResponse
	GetBody() *ModifyShareImageResponseBody
}

type ModifyShareImageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyShareImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyShareImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyShareImageResponse) GoString() string {
	return s.String()
}

func (s *ModifyShareImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyShareImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyShareImageResponse) GetBody() *ModifyShareImageResponseBody {
	return s.Body
}

func (s *ModifyShareImageResponse) SetHeaders(v map[string]*string) *ModifyShareImageResponse {
	s.Headers = v
	return s
}

func (s *ModifyShareImageResponse) SetStatusCode(v int32) *ModifyShareImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyShareImageResponse) SetBody(v *ModifyShareImageResponseBody) *ModifyShareImageResponse {
	s.Body = v
	return s
}

func (s *ModifyShareImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
