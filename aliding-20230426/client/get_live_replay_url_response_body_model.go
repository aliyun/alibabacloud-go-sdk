// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplayUrl(v string) *GetLiveReplayUrlResponseBody
	GetReplayUrl() *string
	SetRequestId(v string) *GetLiveReplayUrlResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetLiveReplayUrlResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetLiveReplayUrlResponseBody
	GetVendorType() *string
}

type GetLiveReplayUrlResponseBody struct {
	// example:
	//
	// https://xxxxxxxx
	ReplayUrl *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetLiveReplayUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlResponseBody) GetReplayUrl() *string {
	return s.ReplayUrl
}

func (s *GetLiveReplayUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveReplayUrlResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetLiveReplayUrlResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetLiveReplayUrlResponseBody) SetReplayUrl(v string) *GetLiveReplayUrlResponseBody {
	s.ReplayUrl = &v
	return s
}

func (s *GetLiveReplayUrlResponseBody) SetRequestId(v string) *GetLiveReplayUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveReplayUrlResponseBody) SetVendorRequestId(v string) *GetLiveReplayUrlResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetLiveReplayUrlResponseBody) SetVendorType(v string) *GetLiveReplayUrlResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetLiveReplayUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
