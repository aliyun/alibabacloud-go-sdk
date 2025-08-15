// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGlobalEventsStorageRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStorageRegion(v string) *UpdateGlobalEventsStorageRegionRequest
	GetStorageRegion() *string
}

type UpdateGlobalEventsStorageRegionRequest struct {
	// The region where you want to store global events.
	//
	// Valid values:
	//
	// 	- ap-southeast-1
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the Singapore region
	//
	//     <!-- -->
	//
	// 	- cn-hangzhou
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the China (Hangzhou) region
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s UpdateGlobalEventsStorageRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *UpdateGlobalEventsStorageRegionRequest) SetStorageRegion(v string) *UpdateGlobalEventsStorageRegionRequest {
	s.StorageRegion = &v
	return s
}

func (s *UpdateGlobalEventsStorageRegionRequest) Validate() error {
	return dara.Validate(s)
}
