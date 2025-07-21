// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppOtaVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionUidList(v []*string) *DeleteAppOtaVersionsRequest
	GetVersionUidList() []*string
}

type DeleteAppOtaVersionsRequest struct {
	VersionUidList []*string `json:"VersionUidList,omitempty" xml:"VersionUidList,omitempty" type:"Repeated"`
}

func (s DeleteAppOtaVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppOtaVersionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsRequest) GetVersionUidList() []*string {
	return s.VersionUidList
}

func (s *DeleteAppOtaVersionsRequest) SetVersionUidList(v []*string) *DeleteAppOtaVersionsRequest {
	s.VersionUidList = v
	return s
}

func (s *DeleteAppOtaVersionsRequest) Validate() error {
	return dara.Validate(s)
}
