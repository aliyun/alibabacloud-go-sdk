// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNetworkInterfaceToInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddNetworkInterfaceToInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddNetworkInterfaceToInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AddNetworkInterfaceToInstanceResponseBody) *AddNetworkInterfaceToInstanceResponse
	GetBody() *AddNetworkInterfaceToInstanceResponseBody
}

type AddNetworkInterfaceToInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddNetworkInterfaceToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddNetworkInterfaceToInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddNetworkInterfaceToInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddNetworkInterfaceToInstanceResponse) GetBody() *AddNetworkInterfaceToInstanceResponseBody {
	return s.Body
}

func (s *AddNetworkInterfaceToInstanceResponse) SetHeaders(v map[string]*string) *AddNetworkInterfaceToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddNetworkInterfaceToInstanceResponse) SetStatusCode(v int32) *AddNetworkInterfaceToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceResponse) SetBody(v *AddNetworkInterfaceToInstanceResponseBody) *AddNetworkInterfaceToInstanceResponse {
	s.Body = v
	return s
}

func (s *AddNetworkInterfaceToInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
