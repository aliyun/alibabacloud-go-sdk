// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListLiveRealtimeLogDeliveryResponseBodyContent) *ListLiveRealtimeLogDeliveryResponseBody
	GetContent() *ListLiveRealtimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *ListLiveRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type ListLiveRealtimeLogDeliveryResponseBody struct {
	// The configurations of real-time log delivery.
	Content *ListLiveRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30559C03-86C9-4EEC-B840-0DC5F5A2189B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponseBody) GetContent() *ListLiveRealtimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *ListLiveRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRealtimeLogDeliveryResponseBody) SetContent(v *ListLiveRealtimeLogDeliveryResponseBodyContent) *ListLiveRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ListLiveRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryResponseBodyContent struct {
	RealtimeLogDeliveryInfo []*ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo `json:"RealtimeLogDeliveryInfo,omitempty" xml:"RealtimeLogDeliveryInfo,omitempty" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContent) GetRealtimeLogDeliveryInfo() []*ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	return s.RealtimeLogDeliveryInfo
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContent) SetRealtimeLogDeliveryInfo(v []*ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) *ListLiveRealtimeLogDeliveryResponseBodyContent {
	s.RealtimeLogDeliveryInfo = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo struct {
	// The ID of the domain name.
	//
	// example:
	//
	// 1001010
	DmId *int32 `json:"DmId,omitempty" xml:"DmId,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the Logstore to which log entries are delivered.
	//
	// example:
	//
	// logstore_example
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// project_example
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The status of real-time log delivery. Valid values:
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDmId() *int32 {
	return s.DmId
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetLogstore() *string {
	return s.Logstore
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetProject() *string {
	return s.Project
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetRegion() *string {
	return s.Region
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetStatus() *string {
	return s.Status
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDmId(v int32) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.DmId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDomainName(v string) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.DomainName = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetLogstore(v string) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetProject(v string) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetRegion(v string) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Region = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetStatus(v string) *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Status = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) Validate() error {
	return dara.Validate(s)
}
