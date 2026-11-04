// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtmpKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRtmpKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRtmpKeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRtmpKeyResponseBody) *UpdateRtmpKeyResponse
	GetBody() *UpdateRtmpKeyResponseBody
}

type UpdateRtmpKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtmpKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRtmpKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRtmpKeyResponse) GetBody() *UpdateRtmpKeyResponseBody {
	return s.Body
}

func (s *UpdateRtmpKeyResponse) SetHeaders(v map[string]*string) *UpdateRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtmpKeyResponse) SetStatusCode(v int32) *UpdateRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtmpKeyResponse) SetBody(v *UpdateRtmpKeyResponseBody) *UpdateRtmpKeyResponse {
	s.Body = v
	return s
}

func (s *UpdateRtmpKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
