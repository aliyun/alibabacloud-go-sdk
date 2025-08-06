// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListVodRealtimeLogDeliveryResponseBodyContent) *ListVodRealtimeLogDeliveryResponseBody
	GetContent() *ListVodRealtimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *ListVodRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type ListVodRealtimeLogDeliveryResponseBody struct {
	Content   *ListVodRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryResponseBody) GetContent() *ListVodRealtimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *ListVodRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodRealtimeLogDeliveryResponseBody) SetContent(v *ListVodRealtimeLogDeliveryResponseBodyContent) *ListVodRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ListVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryResponseBodyContent struct {
	RealtimeLogDeliveryInfo []*ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo `json:"RealtimeLogDeliveryInfo,omitempty" xml:"RealtimeLogDeliveryInfo,omitempty" type:"Repeated"`
}

func (s ListVodRealtimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContent) GetRealtimeLogDeliveryInfo() []*ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	return s.RealtimeLogDeliveryInfo
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContent) SetRealtimeLogDeliveryInfo(v []*ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) *ListVodRealtimeLogDeliveryResponseBodyContent {
	s.RealtimeLogDeliveryInfo = v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo struct {
	DmId       *int32  `json:"DmId,omitempty" xml:"DmId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDmId() *int32 {
	return s.DmId
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetLogstore() *string {
	return s.Logstore
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetProject() *string {
	return s.Project
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetRegion() *string {
	return s.Region
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) GetStatus() *string {
	return s.Status
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDmId(v int32) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.DmId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetDomainName(v string) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.DomainName = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetLogstore(v string) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Logstore = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetProject(v string) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Project = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetRegion(v string) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Region = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) SetStatus(v string) *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo {
	s.Status = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryResponseBodyContentRealtimeLogDeliveryInfo) Validate() error {
	return dara.Validate(s)
}
