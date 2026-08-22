// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchWhitelistGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSearchWhitelistGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSearchWhitelistGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSearchWhitelistGroupResponseBody) *CreateOpenSearchWhitelistGroupResponse
	GetBody() *CreateOpenSearchWhitelistGroupResponseBody
}

type CreateOpenSearchWhitelistGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSearchWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSearchWhitelistGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchWhitelistGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSearchWhitelistGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSearchWhitelistGroupResponse) GetBody() *CreateOpenSearchWhitelistGroupResponseBody {
	return s.Body
}

func (s *CreateOpenSearchWhitelistGroupResponse) SetHeaders(v map[string]*string) *CreateOpenSearchWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponse) SetStatusCode(v int32) *CreateOpenSearchWhitelistGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponse) SetBody(v *CreateOpenSearchWhitelistGroupResponseBody) *CreateOpenSearchWhitelistGroupResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
