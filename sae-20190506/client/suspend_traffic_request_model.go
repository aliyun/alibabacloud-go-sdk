// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SuspendTrafficRequest
	GetAppId() *string
	SetInstanceIds(v string) *SuspendTrafficRequest
	GetInstanceIds() *string
}

type SuspendTrafficRequest struct {
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// c-668727a8-17d86664-41e5bb******,c-668727a8-17d86664-7e4958******
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s SuspendTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendTrafficRequest) GoString() string {
	return s.String()
}

func (s *SuspendTrafficRequest) GetAppId() *string {
	return s.AppId
}

func (s *SuspendTrafficRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *SuspendTrafficRequest) SetAppId(v string) *SuspendTrafficRequest {
	s.AppId = &v
	return s
}

func (s *SuspendTrafficRequest) SetInstanceIds(v string) *SuspendTrafficRequest {
	s.InstanceIds = &v
	return s
}

func (s *SuspendTrafficRequest) Validate() error {
	return dara.Validate(s)
}
