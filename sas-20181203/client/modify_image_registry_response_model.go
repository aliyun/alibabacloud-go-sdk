// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageRegistryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageRegistryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageRegistryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageRegistryResponseBody) *ModifyImageRegistryResponse
	GetBody() *ModifyImageRegistryResponseBody
}

type ModifyImageRegistryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageRegistryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageRegistryResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageRegistryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageRegistryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageRegistryResponse) GetBody() *ModifyImageRegistryResponseBody {
	return s.Body
}

func (s *ModifyImageRegistryResponse) SetHeaders(v map[string]*string) *ModifyImageRegistryResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageRegistryResponse) SetStatusCode(v int32) *ModifyImageRegistryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageRegistryResponse) SetBody(v *ModifyImageRegistryResponseBody) *ModifyImageRegistryResponse {
	s.Body = v
	return s
}

func (s *ModifyImageRegistryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
