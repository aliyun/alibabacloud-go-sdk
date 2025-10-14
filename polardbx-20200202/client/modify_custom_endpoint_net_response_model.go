// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointNetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomEndpointNetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomEndpointNetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomEndpointNetResponseBody) *ModifyCustomEndpointNetResponse
	GetBody() *ModifyCustomEndpointNetResponseBody
}

type ModifyCustomEndpointNetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomEndpointNetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomEndpointNetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointNetResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointNetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomEndpointNetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomEndpointNetResponse) GetBody() *ModifyCustomEndpointNetResponseBody {
	return s.Body
}

func (s *ModifyCustomEndpointNetResponse) SetHeaders(v map[string]*string) *ModifyCustomEndpointNetResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomEndpointNetResponse) SetStatusCode(v int32) *ModifyCustomEndpointNetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomEndpointNetResponse) SetBody(v *ModifyCustomEndpointNetResponseBody) *ModifyCustomEndpointNetResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomEndpointNetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
