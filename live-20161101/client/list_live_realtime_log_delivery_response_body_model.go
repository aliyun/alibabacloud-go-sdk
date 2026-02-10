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
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type ListLiveRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo struct {
	DmId       *int32  `json:"DmId,omitempty" xml:"DmId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
