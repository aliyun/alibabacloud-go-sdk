// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ResumeTrafficRequest
	GetAppId() *string
	SetInstanceIds(v string) *ResumeTrafficRequest
	GetInstanceIds() *string
}

type ResumeTrafficRequest struct {
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// c-668727a8-17d86664-41e5bb******,c-668727a8-17d86664-7e4958******
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s ResumeTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficRequest) GoString() string {
	return s.String()
}

func (s *ResumeTrafficRequest) GetAppId() *string {
	return s.AppId
}

func (s *ResumeTrafficRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ResumeTrafficRequest) SetAppId(v string) *ResumeTrafficRequest {
	s.AppId = &v
	return s
}

func (s *ResumeTrafficRequest) SetInstanceIds(v string) *ResumeTrafficRequest {
	s.InstanceIds = &v
	return s
}

func (s *ResumeTrafficRequest) Validate() error {
	return dara.Validate(s)
}
