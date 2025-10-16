// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDesktopsInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDesktopsInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDesktopsInGroupResponse
	GetStatusCode() *int32
	SetBody(v *DisableDesktopsInGroupResponseBody) *DisableDesktopsInGroupResponse
	GetBody() *DisableDesktopsInGroupResponseBody
}

type DisableDesktopsInGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDesktopsInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDesktopsInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDesktopsInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDesktopsInGroupResponse) GetBody() *DisableDesktopsInGroupResponseBody {
	return s.Body
}

func (s *DisableDesktopsInGroupResponse) SetHeaders(v map[string]*string) *DisableDesktopsInGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableDesktopsInGroupResponse) SetStatusCode(v int32) *DisableDesktopsInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDesktopsInGroupResponse) SetBody(v *DisableDesktopsInGroupResponseBody) *DisableDesktopsInGroupResponse {
	s.Body = v
	return s
}

func (s *DisableDesktopsInGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
