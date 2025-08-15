// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalEventsStorageRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetGlobalEventsStorageRegionResponseBody
	GetRequestId() *string
	SetStorageRegion(v string) *GetGlobalEventsStorageRegionResponseBody
	GetStorageRegion() *string
}

type GetGlobalEventsStorageRegionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0474CD9D-DF37-55D4-8383-D978CFBE13A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region where global events are stored.
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
	// example:
	//
	// cn-hangzhou
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s GetGlobalEventsStorageRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalEventsStorageRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalEventsStorageRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGlobalEventsStorageRegionResponseBody) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *GetGlobalEventsStorageRegionResponseBody) SetRequestId(v string) *GetGlobalEventsStorageRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalEventsStorageRegionResponseBody) SetStorageRegion(v string) *GetGlobalEventsStorageRegionResponseBody {
	s.StorageRegion = &v
	return s
}

func (s *GetGlobalEventsStorageRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
