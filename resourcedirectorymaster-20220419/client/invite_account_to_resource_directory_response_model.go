// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteAccountToResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InviteAccountToResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InviteAccountToResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *InviteAccountToResourceDirectoryResponseBody) *InviteAccountToResourceDirectoryResponse
	GetBody() *InviteAccountToResourceDirectoryResponseBody
}

type InviteAccountToResourceDirectoryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteAccountToResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteAccountToResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InviteAccountToResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InviteAccountToResourceDirectoryResponse) GetBody() *InviteAccountToResourceDirectoryResponseBody {
	return s.Body
}

func (s *InviteAccountToResourceDirectoryResponse) SetHeaders(v map[string]*string) *InviteAccountToResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetStatusCode(v int32) *InviteAccountToResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetBody(v *InviteAccountToResourceDirectoryResponseBody) *InviteAccountToResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
