// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecStatusResponseBody) *ModifyApisecStatusResponse
	GetBody() *ModifyApisecStatusResponseBody
}

type ModifyApisecStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecStatusResponse) GetBody() *ModifyApisecStatusResponseBody {
	return s.Body
}

func (s *ModifyApisecStatusResponse) SetHeaders(v map[string]*string) *ModifyApisecStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecStatusResponse) SetStatusCode(v int32) *ModifyApisecStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecStatusResponse) SetBody(v *ModifyApisecStatusResponseBody) *ModifyApisecStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
