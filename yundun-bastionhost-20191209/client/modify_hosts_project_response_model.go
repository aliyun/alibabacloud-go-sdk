// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostsProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostsProjectResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostsProjectResponseBody) *ModifyHostsProjectResponse
	GetBody() *ModifyHostsProjectResponseBody
}

type ModifyHostsProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostsProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostsProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostsProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostsProjectResponse) GetBody() *ModifyHostsProjectResponseBody {
	return s.Body
}

func (s *ModifyHostsProjectResponse) SetHeaders(v map[string]*string) *ModifyHostsProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostsProjectResponse) SetStatusCode(v int32) *ModifyHostsProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostsProjectResponse) SetBody(v *ModifyHostsProjectResponseBody) *ModifyHostsProjectResponse {
	s.Body = v
	return s
}

func (s *ModifyHostsProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
