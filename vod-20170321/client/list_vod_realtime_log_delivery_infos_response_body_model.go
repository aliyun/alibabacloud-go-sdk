// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListVodRealtimeLogDeliveryInfosResponseBodyContent) *ListVodRealtimeLogDeliveryInfosResponseBody
	GetContent() *ListVodRealtimeLogDeliveryInfosResponseBodyContent
	SetRequestId(v string) *ListVodRealtimeLogDeliveryInfosResponseBody
	GetRequestId() *string
}

type ListVodRealtimeLogDeliveryInfosResponseBody struct {
	Content   *ListVodRealtimeLogDeliveryInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) GetContent() *ListVodRealtimeLogDeliveryInfosResponseBodyContent {
	return s.Content
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) SetContent(v *ListVodRealtimeLogDeliveryInfosResponseBodyContent) *ListVodRealtimeLogDeliveryInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) SetRequestId(v string) *ListVodRealtimeLogDeliveryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryInfosResponseBodyContent struct {
	RealtimeLogDeliveryInfos []*ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" type:"Repeated"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContent) GetRealtimeLogDeliveryInfos() []*ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	return s.RealtimeLogDeliveryInfos
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContent) SetRealtimeLogDeliveryInfos(v []*ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) *ListVodRealtimeLogDeliveryInfosResponseBodyContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetLogstore() *string {
	return s.Logstore
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetProject() *string {
	return s.Project
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetRegion() *string {
	return s.Region
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetProject(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) Validate() error {
	return dara.Validate(s)
}
