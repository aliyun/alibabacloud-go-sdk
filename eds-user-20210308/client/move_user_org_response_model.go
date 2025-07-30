// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveUserOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveUserOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveUserOrgResponse
	GetStatusCode() *int32
	SetBody(v *MoveUserOrgResponseBody) *MoveUserOrgResponse
	GetBody() *MoveUserOrgResponseBody
}

type MoveUserOrgResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveUserOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveUserOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveUserOrgResponse) GoString() string {
	return s.String()
}

func (s *MoveUserOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveUserOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveUserOrgResponse) GetBody() *MoveUserOrgResponseBody {
	return s.Body
}

func (s *MoveUserOrgResponse) SetHeaders(v map[string]*string) *MoveUserOrgResponse {
	s.Headers = v
	return s
}

func (s *MoveUserOrgResponse) SetStatusCode(v int32) *MoveUserOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveUserOrgResponse) SetBody(v *MoveUserOrgResponseBody) *MoveUserOrgResponse {
	s.Body = v
	return s
}

func (s *MoveUserOrgResponse) Validate() error {
	return dara.Validate(s)
}
