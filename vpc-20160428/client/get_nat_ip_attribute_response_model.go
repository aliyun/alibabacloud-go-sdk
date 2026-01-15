// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNatIpAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNatIpAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetNatIpAttributeResponseBody) *GetNatIpAttributeResponse
	GetBody() *GetNatIpAttributeResponseBody
}

type GetNatIpAttributeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatIpAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatIpAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetNatIpAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNatIpAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNatIpAttributeResponse) GetBody() *GetNatIpAttributeResponseBody {
	return s.Body
}

func (s *GetNatIpAttributeResponse) SetHeaders(v map[string]*string) *GetNatIpAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetNatIpAttributeResponse) SetStatusCode(v int32) *GetNatIpAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatIpAttributeResponse) SetBody(v *GetNatIpAttributeResponseBody) *GetNatIpAttributeResponse {
	s.Body = v
	return s
}

func (s *GetNatIpAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
