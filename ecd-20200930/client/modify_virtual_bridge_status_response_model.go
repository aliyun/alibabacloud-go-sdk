// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVirtualBridgeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVirtualBridgeStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVirtualBridgeStatusResponseBody) *ModifyVirtualBridgeStatusResponse
	GetBody() *ModifyVirtualBridgeStatusResponseBody
}

type ModifyVirtualBridgeStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVirtualBridgeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVirtualBridgeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVirtualBridgeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVirtualBridgeStatusResponse) GetBody() *ModifyVirtualBridgeStatusResponseBody {
	return s.Body
}

func (s *ModifyVirtualBridgeStatusResponse) SetHeaders(v map[string]*string) *ModifyVirtualBridgeStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVirtualBridgeStatusResponse) SetStatusCode(v int32) *ModifyVirtualBridgeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVirtualBridgeStatusResponse) SetBody(v *ModifyVirtualBridgeStatusResponseBody) *ModifyVirtualBridgeStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyVirtualBridgeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
