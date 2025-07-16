// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListRealtimeLogDeliveryInfosResponseBodyContent) *ListRealtimeLogDeliveryInfosResponseBody
	GetContent() *ListRealtimeLogDeliveryInfosResponseBodyContent
	SetRequestId(v string) *ListRealtimeLogDeliveryInfosResponseBody
	GetRequestId() *string
}

type ListRealtimeLogDeliveryInfosResponseBody struct {
	// The information about real-time log delivery.
	Content *ListRealtimeLogDeliveryInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 95D5B69F-8AEC-419B-8F3A-612B35032B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeLogDeliveryInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) GetContent() *ListRealtimeLogDeliveryInfosResponseBodyContent {
	return s.Content
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) SetContent(v *ListRealtimeLogDeliveryInfosResponseBodyContent) *ListRealtimeLogDeliveryInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) SetRequestId(v string) *ListRealtimeLogDeliveryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeLogDeliveryInfosResponseBodyContent struct {
	RealtimeLogDeliveryInfos []*ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" type:"Repeated"`
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContent) GetRealtimeLogDeliveryInfos() []*ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	return s.RealtimeLogDeliveryInfos
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContent) SetRealtimeLogDeliveryInfos(v []*ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) *ListRealtimeLogDeliveryInfosResponseBodyContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos struct {
	// The name of the Logstore that collects log data from Alibaba Cloud CDN in real time.
	//
	// example:
	//
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed. For more information, see [Regions that support real-time log delivery](https://help.aliyun.com/document_detail/144883.html).
	//
	// example:
	//
	// ch-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetLogstore() *string {
	return s.Logstore
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetProject() *string {
	return s.Project
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GetRegion() *string {
	return s.Region
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetProject(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) Validate() error {
	return dara.Validate(s)
}
