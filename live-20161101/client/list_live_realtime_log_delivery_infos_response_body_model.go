// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListLiveRealtimeLogDeliveryInfosResponseBodyContent) *ListLiveRealtimeLogDeliveryInfosResponseBody
	GetContent() *ListLiveRealtimeLogDeliveryInfosResponseBodyContent
	SetRequestId(v string) *ListLiveRealtimeLogDeliveryInfosResponseBody
	GetRequestId() *string
}

type ListLiveRealtimeLogDeliveryInfosResponseBody struct {
	// Details about the configuration of real-time log delivery.
	Content *ListLiveRealtimeLogDeliveryInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 95D5B69F-8AEC-419B-8F3A-612B35032B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBody) GetContent() *ListLiveRealtimeLogDeliveryInfosResponseBodyContent {
	return s.Content
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBody) SetContent(v *ListLiveRealtimeLogDeliveryInfosResponseBodyContent) *ListLiveRealtimeLogDeliveryInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBody) SetRequestId(v string) *ListLiveRealtimeLogDeliveryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryInfosResponseBodyContent struct {
	RealtimeLogDeliveryInfos []*ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContent) GetRealtimeLogDeliveryInfos() []*ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	return s.RealtimeLogDeliveryInfos
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContent) SetRealtimeLogDeliveryInfos(v []*ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) *ListLiveRealtimeLogDeliveryInfosResponseBodyContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos struct {
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
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetLogstore() *string {
	return s.Logstore
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetProject() *string {
	return s.Project
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetRegion() *string {
	return s.Region
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetProject(v string) *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) Validate() error {
	return dara.Validate(s)
}
