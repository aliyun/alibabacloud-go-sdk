// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchWhitelistGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOpenSearchWhitelistGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOpenSearchWhitelistGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOpenSearchWhitelistGroupResponseBody) *ModifyOpenSearchWhitelistGroupResponse
	GetBody() *ModifyOpenSearchWhitelistGroupResponseBody
}

type ModifyOpenSearchWhitelistGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOpenSearchWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOpenSearchWhitelistGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchWhitelistGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOpenSearchWhitelistGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOpenSearchWhitelistGroupResponse) GetBody() *ModifyOpenSearchWhitelistGroupResponseBody {
	return s.Body
}

func (s *ModifyOpenSearchWhitelistGroupResponse) SetHeaders(v map[string]*string) *ModifyOpenSearchWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponse) SetStatusCode(v int32) *ModifyOpenSearchWhitelistGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponse) SetBody(v *ModifyOpenSearchWhitelistGroupResponseBody) *ModifyOpenSearchWhitelistGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
