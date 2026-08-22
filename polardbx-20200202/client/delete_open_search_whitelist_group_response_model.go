// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchWhitelistGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpenSearchWhitelistGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpenSearchWhitelistGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpenSearchWhitelistGroupResponseBody) *DeleteOpenSearchWhitelistGroupResponse
	GetBody() *DeleteOpenSearchWhitelistGroupResponseBody
}

type DeleteOpenSearchWhitelistGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpenSearchWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpenSearchWhitelistGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchWhitelistGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpenSearchWhitelistGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpenSearchWhitelistGroupResponse) GetBody() *DeleteOpenSearchWhitelistGroupResponseBody {
	return s.Body
}

func (s *DeleteOpenSearchWhitelistGroupResponse) SetHeaders(v map[string]*string) *DeleteOpenSearchWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponse) SetStatusCode(v int32) *DeleteOpenSearchWhitelistGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponse) SetBody(v *DeleteOpenSearchWhitelistGroupResponseBody) *DeleteOpenSearchWhitelistGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
