// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcVSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserVpcVSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserVpcVSwitchResponse
	GetStatusCode() *int32
	SetBody(v *GetUserVpcVSwitchResponseBody) *GetUserVpcVSwitchResponse
	GetBody() *GetUserVpcVSwitchResponseBody
}

type GetUserVpcVSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserVpcVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserVpcVSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcVSwitchResponse) GoString() string {
	return s.String()
}

func (s *GetUserVpcVSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserVpcVSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserVpcVSwitchResponse) GetBody() *GetUserVpcVSwitchResponseBody {
	return s.Body
}

func (s *GetUserVpcVSwitchResponse) SetHeaders(v map[string]*string) *GetUserVpcVSwitchResponse {
	s.Headers = v
	return s
}

func (s *GetUserVpcVSwitchResponse) SetStatusCode(v int32) *GetUserVpcVSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserVpcVSwitchResponse) SetBody(v *GetUserVpcVSwitchResponseBody) *GetUserVpcVSwitchResponse {
	s.Body = v
	return s
}

func (s *GetUserVpcVSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
