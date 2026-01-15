// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpCidrAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNatIpCidrAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNatIpCidrAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetNatIpCidrAttributeResponseBody) *GetNatIpCidrAttributeResponse
	GetBody() *GetNatIpCidrAttributeResponseBody
}

type GetNatIpCidrAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatIpCidrAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatIpCidrAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpCidrAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetNatIpCidrAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNatIpCidrAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNatIpCidrAttributeResponse) GetBody() *GetNatIpCidrAttributeResponseBody {
	return s.Body
}

func (s *GetNatIpCidrAttributeResponse) SetHeaders(v map[string]*string) *GetNatIpCidrAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetNatIpCidrAttributeResponse) SetStatusCode(v int32) *GetNatIpCidrAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatIpCidrAttributeResponse) SetBody(v *GetNatIpCidrAttributeResponseBody) *GetNatIpCidrAttributeResponse {
	s.Body = v
	return s
}

func (s *GetNatIpCidrAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
