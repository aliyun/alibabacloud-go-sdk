// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserGroupInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserGroupInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserGroupInfoResponseBody) *GetUserGroupInfoResponse
	GetBody() *GetUserGroupInfoResponseBody
}

type GetUserGroupInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGroupInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserGroupInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserGroupInfoResponse) GetBody() *GetUserGroupInfoResponseBody {
	return s.Body
}

func (s *GetUserGroupInfoResponse) SetHeaders(v map[string]*string) *GetUserGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupInfoResponse) SetStatusCode(v int32) *GetUserGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupInfoResponse) SetBody(v *GetUserGroupInfoResponseBody) *GetUserGroupInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserGroupInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
