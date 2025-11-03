// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpCidrAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatIpCidrAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatIpCidrAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatIpCidrAttributeResponseBody) *ModifyNatIpCidrAttributeResponse
	GetBody() *ModifyNatIpCidrAttributeResponseBody
}

type ModifyNatIpCidrAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatIpCidrAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatIpCidrAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpCidrAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatIpCidrAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatIpCidrAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatIpCidrAttributeResponse) GetBody() *ModifyNatIpCidrAttributeResponseBody {
	return s.Body
}

func (s *ModifyNatIpCidrAttributeResponse) SetHeaders(v map[string]*string) *ModifyNatIpCidrAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatIpCidrAttributeResponse) SetStatusCode(v int32) *ModifyNatIpCidrAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatIpCidrAttributeResponse) SetBody(v *ModifyNatIpCidrAttributeResponseBody) *ModifyNatIpCidrAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNatIpCidrAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
