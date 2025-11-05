// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListRealtimeLogDeliveryResponseBodyContent) *ListRealtimeLogDeliveryResponseBody
	GetContent() *ListRealtimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *ListRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type ListRealtimeLogDeliveryResponseBody struct {
	// The logging information.
	Content *ListRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 30559C03-86C9-4EEC-B840-0DC5F5A2189B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryResponseBody) GetContent() *ListRealtimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *ListRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeLogDeliveryResponseBody) SetContent(v *ListRealtimeLogDeliveryResponseBodyContent) *ListRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ListRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRealtimeLogDeliveryResponseBodyContent struct {
	RealtimeLogDeliveryInfo []*ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo `json:"RealtimeLogDeliveryInfo,omitempty" xml:"RealtimeLogDeliveryInfo,omitempty" type:"Repeated"`
}

func (s ListRealtimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryResponseBodyContent) GetRealtimeLogDeliveryInfo() []*ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	return s.RealtimeLogDeliveryInfo
}

func (s *ListRealtimeLogDeliveryResponseBodyContent) SetRealtimeLogDeliveryInfo(v []*ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) *ListRealtimeLogDeliveryResponseBodyContent {
	s.RealtimeLogDeliveryInfo = v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContent) Validate() error {
	if s.RealtimeLogDeliveryInfo != nil {
		for _, item := range s.RealtimeLogDeliveryInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo struct {
	// The domain ID.
	//
	// example:
	//
	// 1001010
	DmId *int32 `json:"DmId,omitempty" xml:"DmId,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the Logstore where log entries are stored.
	//
	// example:
	//
	// test
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// test
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// example:
	//
	// cn-hangzhou-corp
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The status of real-time log delivery.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDmId() *int32 {
	return s.DmId
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDomain() *string {
	return s.Domain
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetLogstore() *string {
	return s.Logstore
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetProject() *string {
	return s.Project
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetRegion() *string {
	return s.Region
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetStatus() *string {
	return s.Status
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDmId(v int32) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.DmId = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDomain(v string) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Domain = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetLogstore(v string) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Logstore = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetProject(v string) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Project = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetRegion(v string) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Region = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetStatus(v string) *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Status = &v
	return s
}

func (s *ListRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) Validate() error {
	return dara.Validate(s)
}
