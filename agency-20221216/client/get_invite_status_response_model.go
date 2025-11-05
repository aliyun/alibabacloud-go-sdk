// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInviteStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInviteStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInviteStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetInviteStatusResponseBody) *GetInviteStatusResponse
	GetBody() *GetInviteStatusResponseBody
}

type GetInviteStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInviteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInviteStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusResponse) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInviteStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInviteStatusResponse) GetBody() *GetInviteStatusResponseBody {
	return s.Body
}

func (s *GetInviteStatusResponse) SetHeaders(v map[string]*string) *GetInviteStatusResponse {
	s.Headers = v
	return s
}

func (s *GetInviteStatusResponse) SetStatusCode(v int32) *GetInviteStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInviteStatusResponse) SetBody(v *GetInviteStatusResponseBody) *GetInviteStatusResponse {
	s.Body = v
	return s
}

func (s *GetInviteStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
