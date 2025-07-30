// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveOrgResponse
	GetStatusCode() *int32
	SetBody(v *MoveOrgResponseBody) *MoveOrgResponse
	GetBody() *MoveOrgResponseBody
}

type MoveOrgResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveOrgResponse) GoString() string {
	return s.String()
}

func (s *MoveOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveOrgResponse) GetBody() *MoveOrgResponseBody {
	return s.Body
}

func (s *MoveOrgResponse) SetHeaders(v map[string]*string) *MoveOrgResponse {
	s.Headers = v
	return s
}

func (s *MoveOrgResponse) SetStatusCode(v int32) *MoveOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveOrgResponse) SetBody(v *MoveOrgResponseBody) *MoveOrgResponse {
	s.Body = v
	return s
}

func (s *MoveOrgResponse) Validate() error {
	return dara.Validate(s)
}
