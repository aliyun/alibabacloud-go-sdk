// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpsetsBandwidthLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthAllocationType(v string) *GetIpsetsBandwidthLimitResponseBody
	GetBandwidthAllocationType() *string
	SetBandwidthLimit(v int64) *GetIpsetsBandwidthLimitResponseBody
	GetBandwidthLimit() *int64
	SetRequestId(v string) *GetIpsetsBandwidthLimitResponseBody
	GetRequestId() *string
}

type GetIpsetsBandwidthLimitResponseBody struct {
	// The type of the bandwidth that is allocated.
	//
	// 	- **ShareBandwidth:*	- shared bandwidth.
	//
	// 	- **ExclusiveBandwidth:*	- dedicated bandwidth.
	//
	// example:
	//
	// ShareBandwidth
	BandwidthAllocationType *string `json:"BandwidthAllocationType,omitempty" xml:"BandwidthAllocationType,omitempty"`
	// The maximum bandwidth of the acceleration area. Unit: Mbit/s.
	//
	// example:
	//
	// 20
	BandwidthLimit *int64 `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpsetsBandwidthLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpsetsBandwidthLimitResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpsetsBandwidthLimitResponseBody) GetBandwidthAllocationType() *string {
	return s.BandwidthAllocationType
}

func (s *GetIpsetsBandwidthLimitResponseBody) GetBandwidthLimit() *int64 {
	return s.BandwidthLimit
}

func (s *GetIpsetsBandwidthLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpsetsBandwidthLimitResponseBody) SetBandwidthAllocationType(v string) *GetIpsetsBandwidthLimitResponseBody {
	s.BandwidthAllocationType = &v
	return s
}

func (s *GetIpsetsBandwidthLimitResponseBody) SetBandwidthLimit(v int64) *GetIpsetsBandwidthLimitResponseBody {
	s.BandwidthLimit = &v
	return s
}

func (s *GetIpsetsBandwidthLimitResponseBody) SetRequestId(v string) *GetIpsetsBandwidthLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpsetsBandwidthLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
