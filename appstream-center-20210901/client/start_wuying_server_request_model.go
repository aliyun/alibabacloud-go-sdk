// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWuyingServerIdList(v []*string) *StartWuyingServerRequest
	GetWuyingServerIdList() []*string
}

type StartWuyingServerRequest struct {
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
}

func (s StartWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *StartWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *StartWuyingServerRequest) SetWuyingServerIdList(v []*string) *StartWuyingServerRequest {
	s.WuyingServerIdList = v
	return s
}

func (s *StartWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
