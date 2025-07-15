// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *DeleteAppsRequest
	GetAppIdList() []*string
}

type DeleteAppsRequest struct {
	// The IDs of the applications.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
}

func (s DeleteAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppsRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *DeleteAppsRequest) SetAppIdList(v []*string) *DeleteAppsRequest {
	s.AppIdList = v
	return s
}

func (s *DeleteAppsRequest) Validate() error {
	return dara.Validate(s)
}
