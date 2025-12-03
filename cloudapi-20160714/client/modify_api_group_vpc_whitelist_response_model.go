// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupVpcWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiGroupVpcWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiGroupVpcWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiGroupVpcWhitelistResponseBody) *ModifyApiGroupVpcWhitelistResponse
	GetBody() *ModifyApiGroupVpcWhitelistResponseBody
}

type ModifyApiGroupVpcWhitelistResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiGroupVpcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiGroupVpcWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiGroupVpcWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiGroupVpcWhitelistResponse) GetBody() *ModifyApiGroupVpcWhitelistResponseBody {
	return s.Body
}

func (s *ModifyApiGroupVpcWhitelistResponse) SetHeaders(v map[string]*string) *ModifyApiGroupVpcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupVpcWhitelistResponse) SetStatusCode(v int32) *ModifyApiGroupVpcWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistResponse) SetBody(v *ModifyApiGroupVpcWhitelistResponseBody) *ModifyApiGroupVpcWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyApiGroupVpcWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
