// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWuyingServerIdList(v []*string) *RestartWuyingServerRequest
	GetWuyingServerIdList() []*string
}

type RestartWuyingServerRequest struct {
	// The list of workstation IDs.
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
}

func (s RestartWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *RestartWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *RestartWuyingServerRequest) SetWuyingServerIdList(v []*string) *RestartWuyingServerRequest {
	s.WuyingServerIdList = v
	return s
}

func (s *RestartWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
