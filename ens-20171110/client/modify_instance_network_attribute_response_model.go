// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceNetworkAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceNetworkAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceNetworkAttributeResponseBody) *ModifyInstanceNetworkAttributeResponse
	GetBody() *ModifyInstanceNetworkAttributeResponseBody
}

type ModifyInstanceNetworkAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNetworkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNetworkAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceNetworkAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceNetworkAttributeResponse) GetBody() *ModifyInstanceNetworkAttributeResponseBody {
	return s.Body
}

func (s *ModifyInstanceNetworkAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceNetworkAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNetworkAttributeResponse) SetStatusCode(v int32) *ModifyInstanceNetworkAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNetworkAttributeResponse) SetBody(v *ModifyInstanceNetworkAttributeResponseBody) *ModifyInstanceNetworkAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceNetworkAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
