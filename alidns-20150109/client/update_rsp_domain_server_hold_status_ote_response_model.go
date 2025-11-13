// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerHoldStatusOteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRspDomainServerHoldStatusOteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRspDomainServerHoldStatusOteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRspDomainServerHoldStatusOteResponseBody) *UpdateRspDomainServerHoldStatusOteResponse
	GetBody() *UpdateRspDomainServerHoldStatusOteResponseBody
}

type UpdateRspDomainServerHoldStatusOteResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRspDomainServerHoldStatusOteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRspDomainServerHoldStatusOteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerHoldStatusOteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) GetBody() *UpdateRspDomainServerHoldStatusOteResponseBody {
	return s.Body
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) SetHeaders(v map[string]*string) *UpdateRspDomainServerHoldStatusOteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) SetStatusCode(v int32) *UpdateRspDomainServerHoldStatusOteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) SetBody(v *UpdateRspDomainServerHoldStatusOteResponseBody) *UpdateRspDomainServerHoldStatusOteResponse {
	s.Body = v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
