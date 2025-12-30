// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindZoneVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindZoneVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindZoneVpcResponse
	GetStatusCode() *int32
	SetBody(v *BindZoneVpcResponseBody) *BindZoneVpcResponse
	GetBody() *BindZoneVpcResponseBody
}

type BindZoneVpcResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindZoneVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindZoneVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s BindZoneVpcResponse) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindZoneVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindZoneVpcResponse) GetBody() *BindZoneVpcResponseBody {
	return s.Body
}

func (s *BindZoneVpcResponse) SetHeaders(v map[string]*string) *BindZoneVpcResponse {
	s.Headers = v
	return s
}

func (s *BindZoneVpcResponse) SetStatusCode(v int32) *BindZoneVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindZoneVpcResponse) SetBody(v *BindZoneVpcResponseBody) *BindZoneVpcResponse {
	s.Body = v
	return s
}

func (s *BindZoneVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
