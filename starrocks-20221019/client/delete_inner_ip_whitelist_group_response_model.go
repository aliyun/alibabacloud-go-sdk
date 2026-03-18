// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInnerIpWhitelistGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInnerIpWhitelistGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInnerIpWhitelistGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInnerIpWhitelistGroupResponseBody) *DeleteInnerIpWhitelistGroupResponse
	GetBody() *DeleteInnerIpWhitelistGroupResponseBody
}

type DeleteInnerIpWhitelistGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInnerIpWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInnerIpWhitelistGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInnerIpWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteInnerIpWhitelistGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInnerIpWhitelistGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInnerIpWhitelistGroupResponse) GetBody() *DeleteInnerIpWhitelistGroupResponseBody {
	return s.Body
}

func (s *DeleteInnerIpWhitelistGroupResponse) SetHeaders(v map[string]*string) *DeleteInnerIpWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponse) SetStatusCode(v int32) *DeleteInnerIpWhitelistGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponse) SetBody(v *DeleteInnerIpWhitelistGroupResponseBody) *DeleteInnerIpWhitelistGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteInnerIpWhitelistGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
