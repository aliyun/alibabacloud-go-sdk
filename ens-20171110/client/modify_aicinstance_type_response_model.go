// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAICInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAICInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAICInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAICInstanceTypeResponseBody) *ModifyAICInstanceTypeResponse
	GetBody() *ModifyAICInstanceTypeResponseBody
}

type ModifyAICInstanceTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAICInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAICInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAICInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAICInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAICInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAICInstanceTypeResponse) GetBody() *ModifyAICInstanceTypeResponseBody {
	return s.Body
}

func (s *ModifyAICInstanceTypeResponse) SetHeaders(v map[string]*string) *ModifyAICInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAICInstanceTypeResponse) SetStatusCode(v int32) *ModifyAICInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAICInstanceTypeResponse) SetBody(v *ModifyAICInstanceTypeResponseBody) *ModifyAICInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyAICInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
