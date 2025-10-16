// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDesktopOversoldGroupResponseBody) *CreateDesktopOversoldGroupResponse
	GetBody() *CreateDesktopOversoldGroupResponseBody
}

type CreateDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDesktopOversoldGroupResponse) GetBody() *CreateDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *CreateDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *CreateDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopOversoldGroupResponse) SetStatusCode(v int32) *CreateDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDesktopOversoldGroupResponse) SetBody(v *CreateDesktopOversoldGroupResponseBody) *CreateDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDesktopOversoldGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
