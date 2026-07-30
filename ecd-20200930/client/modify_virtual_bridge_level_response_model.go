// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVirtualBridgeLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVirtualBridgeLevelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVirtualBridgeLevelResponseBody) *ModifyVirtualBridgeLevelResponse
	GetBody() *ModifyVirtualBridgeLevelResponseBody
}

type ModifyVirtualBridgeLevelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVirtualBridgeLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVirtualBridgeLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeLevelResponse) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVirtualBridgeLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVirtualBridgeLevelResponse) GetBody() *ModifyVirtualBridgeLevelResponseBody {
	return s.Body
}

func (s *ModifyVirtualBridgeLevelResponse) SetHeaders(v map[string]*string) *ModifyVirtualBridgeLevelResponse {
	s.Headers = v
	return s
}

func (s *ModifyVirtualBridgeLevelResponse) SetStatusCode(v int32) *ModifyVirtualBridgeLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVirtualBridgeLevelResponse) SetBody(v *ModifyVirtualBridgeLevelResponseBody) *ModifyVirtualBridgeLevelResponse {
	s.Body = v
	return s
}

func (s *ModifyVirtualBridgeLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
