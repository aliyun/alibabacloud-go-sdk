// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUpgradableVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryUpgradableVersionsRequest
	GetInstanceId() *string
	SetMinor(v bool) *QueryUpgradableVersionsRequest
	GetMinor() *bool
}

type QueryUpgradableVersionsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query the minor versions that you can upgrade to. Default value: true. Valid values:
	//
	// 	- true: The minor versions that you can upgrade to.
	//
	// 	- false: The major versions that you can upgrade to.
	//
	// example:
	//
	// true
	Minor *bool `json:"Minor,omitempty" xml:"Minor,omitempty"`
}

func (s QueryUpgradableVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUpgradableVersionsRequest) GoString() string {
	return s.String()
}

func (s *QueryUpgradableVersionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryUpgradableVersionsRequest) GetMinor() *bool {
	return s.Minor
}

func (s *QueryUpgradableVersionsRequest) SetInstanceId(v string) *QueryUpgradableVersionsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryUpgradableVersionsRequest) SetMinor(v bool) *QueryUpgradableVersionsRequest {
	s.Minor = &v
	return s
}

func (s *QueryUpgradableVersionsRequest) Validate() error {
	return dara.Validate(s)
}
