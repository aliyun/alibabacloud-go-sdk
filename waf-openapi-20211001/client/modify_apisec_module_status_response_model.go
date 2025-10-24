// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecModuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecModuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecModuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecModuleStatusResponseBody) *ModifyApisecModuleStatusResponse
	GetBody() *ModifyApisecModuleStatusResponseBody
}

type ModifyApisecModuleStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecModuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecModuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecModuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecModuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecModuleStatusResponse) GetBody() *ModifyApisecModuleStatusResponseBody {
	return s.Body
}

func (s *ModifyApisecModuleStatusResponse) SetHeaders(v map[string]*string) *ModifyApisecModuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecModuleStatusResponse) SetStatusCode(v int32) *ModifyApisecModuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecModuleStatusResponse) SetBody(v *ModifyApisecModuleStatusResponseBody) *ModifyApisecModuleStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecModuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
