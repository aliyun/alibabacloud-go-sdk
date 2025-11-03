// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicIpAddressPoolServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublicIpAddressPoolServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublicIpAddressPoolServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPublicIpAddressPoolServiceStatusResponseBody) *GetPublicIpAddressPoolServiceStatusResponse
	GetBody() *GetPublicIpAddressPoolServiceStatusResponseBody
}

type GetPublicIpAddressPoolServiceStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicIpAddressPoolServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicIpAddressPoolServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublicIpAddressPoolServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) GetBody() *GetPublicIpAddressPoolServiceStatusResponseBody {
	return s.Body
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) SetHeaders(v map[string]*string) *GetPublicIpAddressPoolServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) SetStatusCode(v int32) *GetPublicIpAddressPoolServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) SetBody(v *GetPublicIpAddressPoolServiceStatusResponseBody) *GetPublicIpAddressPoolServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
