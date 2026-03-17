// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyACLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyACLResponseBody) *ModifyACLResponse
	GetBody() *ModifyACLResponseBody
}

type ModifyACLResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyACLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLResponse) GoString() string {
	return s.String()
}

func (s *ModifyACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyACLResponse) GetBody() *ModifyACLResponseBody {
	return s.Body
}

func (s *ModifyACLResponse) SetHeaders(v map[string]*string) *ModifyACLResponse {
	s.Headers = v
	return s
}

func (s *ModifyACLResponse) SetStatusCode(v int32) *ModifyACLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyACLResponse) SetBody(v *ModifyACLResponseBody) *ModifyACLResponse {
	s.Body = v
	return s
}

func (s *ModifyACLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
