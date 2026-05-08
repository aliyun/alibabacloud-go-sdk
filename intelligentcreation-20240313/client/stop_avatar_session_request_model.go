// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAvatarSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *StopAvatarSessionRequest
	GetProjectId() *string
	SetSessionId(v string) *StopAvatarSessionRequest
	GetSessionId() *string
}

type StopAvatarSessionRequest struct {
	// example:
	//
	// 124900036
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hojazg0u9lvysk0uyczgd79be2a4hkr9ijrblmb5qohi5iaja3p5j633doqj4t2uu3sek2i49hzkao0bli4bch4tnloyx22odd7sot9dxl5xfd0hbp7fl9dehnqofkb9csebf0nuezj8bwgec8ei6dby0encu5y88ky6oqensuqnj
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s StopAvatarSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAvatarSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StopAvatarSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopAvatarSessionRequest) SetProjectId(v string) *StopAvatarSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAvatarSessionRequest) SetSessionId(v string) *StopAvatarSessionRequest {
	s.SessionId = &v
	return s
}

func (s *StopAvatarSessionRequest) Validate() error {
	return dara.Validate(s)
}
