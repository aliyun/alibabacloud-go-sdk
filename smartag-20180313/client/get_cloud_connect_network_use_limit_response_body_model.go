// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudConnectNetworkUseLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCloudConnectNetworkUseLimitResponseBody
	GetRequestId() *string
	SetTotalAmount(v int32) *GetCloudConnectNetworkUseLimitResponseBody
	GetTotalAmount() *int32
	SetUsedAmount(v int32) *GetCloudConnectNetworkUseLimitResponseBody
	GetUsedAmount() *int32
}

type GetCloudConnectNetworkUseLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BCD04867-56C3-43DC-8FEF-923EFB8B56DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of CCN instances that the current account can create in the specified region.
	//
	// example:
	//
	// 10
	TotalAmount *int32 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The number of CCN instances that you have created.
	//
	// example:
	//
	// 6
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
}

func (s GetCloudConnectNetworkUseLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudConnectNetworkUseLimitResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) GetTotalAmount() *int32 {
	return s.TotalAmount
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) SetRequestId(v string) *GetCloudConnectNetworkUseLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) SetTotalAmount(v int32) *GetCloudConnectNetworkUseLimitResponseBody {
	s.TotalAmount = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) SetUsedAmount(v int32) *GetCloudConnectNetworkUseLimitResponseBody {
	s.UsedAmount = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
