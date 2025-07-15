// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveCenterTransferInfoList(v *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) *DescribeLiveCenterTransferResponseBody
	GetLiveCenterTransferInfoList() *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList
	SetRequestId(v string) *DescribeLiveCenterTransferResponseBody
	GetRequestId() *string
}

type DescribeLiveCenterTransferResponseBody struct {
	// The stream relay information.
	LiveCenterTransferInfoList *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList `json:"LiveCenterTransferInfoList,omitempty" xml:"LiveCenterTransferInfoList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C4865B85-664B-19D3-BB16-C62FB83C8226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveCenterTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterTransferResponseBody) GetLiveCenterTransferInfoList() *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList {
	return s.LiveCenterTransferInfoList
}

func (s *DescribeLiveCenterTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveCenterTransferResponseBody) SetLiveCenterTransferInfoList(v *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) *DescribeLiveCenterTransferResponseBody {
	s.LiveCenterTransferInfoList = v
	return s
}

func (s *DescribeLiveCenterTransferResponseBody) SetRequestId(v string) *DescribeLiveCenterTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList struct {
	LiveCenterTransferInfo []*DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo `json:"LiveCenterTransferInfo,omitempty" xml:"LiveCenterTransferInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) GetLiveCenterTransferInfo() []*DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	return s.LiveCenterTransferInfo
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) SetLiveCenterTransferInfo(v []*DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList {
	s.LiveCenterTransferInfo = v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// teststream
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The third-party URL to which the live stream is relayed.
	//
	// example:
	//
	// rtmp://push.example2.aliyunlive.com/testapp1/teststream2
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of stream relay.
	//
	// example:
	//
	// 2022-04-29T15:16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of stream relay.
	//
	// example:
	//
	// 2022-04-28T15:16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The validity period of stream relay. Valid values:
	//
	// 	- **always**: The stream can always be relayed.
	//
	// 	- **time**: The stream can be relayed in a specified time period.
	//
	// example:
	//
	// time
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetAppName(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetDomainName(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetDstUrl(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.DstUrl = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetEndTime(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetStartTime(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetStreamName(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) SetTransferArgs(v string) *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo {
	s.TransferArgs = &v
	return s
}

func (s *DescribeLiveCenterTransferResponseBodyLiveCenterTransferInfoListLiveCenterTransferInfo) Validate() error {
	return dara.Validate(s)
}
