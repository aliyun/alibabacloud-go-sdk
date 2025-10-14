// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkAttributeResponseBody) *ModifyNetworkAttributeResponse
	GetBody() *ModifyNetworkAttributeResponseBody
}

type ModifyNetworkAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkAttributeResponse) GetBody() *ModifyNetworkAttributeResponseBody {
	return s.Body
}

func (s *ModifyNetworkAttributeResponse) SetHeaders(v map[string]*string) *ModifyNetworkAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkAttributeResponse) SetStatusCode(v int32) *ModifyNetworkAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkAttributeResponse) SetBody(v *ModifyNetworkAttributeResponseBody) *ModifyNetworkAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
