// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudVendorAccountAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthId(v int64) *DeleteCloudVendorAccountAKRequest
	GetAuthId() *int64
	SetAuthModules(v []*string) *DeleteCloudVendorAccountAKRequest
	GetAuthModules() []*string
}

type DeleteCloudVendorAccountAKRequest struct {
	// The unique ID of the AccessKey pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2363
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The modules that are associated with the AccessKey pair.
	AuthModules []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
}

func (s DeleteCloudVendorAccountAKRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudVendorAccountAKRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudVendorAccountAKRequest) GetAuthId() *int64 {
	return s.AuthId
}

func (s *DeleteCloudVendorAccountAKRequest) GetAuthModules() []*string {
	return s.AuthModules
}

func (s *DeleteCloudVendorAccountAKRequest) SetAuthId(v int64) *DeleteCloudVendorAccountAKRequest {
	s.AuthId = &v
	return s
}

func (s *DeleteCloudVendorAccountAKRequest) SetAuthModules(v []*string) *DeleteCloudVendorAccountAKRequest {
	s.AuthModules = v
	return s
}

func (s *DeleteCloudVendorAccountAKRequest) Validate() error {
	return dara.Validate(s)
}
