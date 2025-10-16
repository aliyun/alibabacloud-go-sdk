// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopHostNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopHostNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopHostNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopHostNameResponseBody) *ModifyDesktopHostNameResponse
	GetBody() *ModifyDesktopHostNameResponseBody
}

type ModifyDesktopHostNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopHostNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopHostNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopHostNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopHostNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopHostNameResponse) GetBody() *ModifyDesktopHostNameResponseBody {
	return s.Body
}

func (s *ModifyDesktopHostNameResponse) SetHeaders(v map[string]*string) *ModifyDesktopHostNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopHostNameResponse) SetStatusCode(v int32) *ModifyDesktopHostNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopHostNameResponse) SetBody(v *ModifyDesktopHostNameResponseBody) *ModifyDesktopHostNameResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopHostNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
