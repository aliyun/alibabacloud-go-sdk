// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeSystemEventMetaListResponseBody
	GetCode() *int32
	SetData(v *DescribeSystemEventMetaListResponseBodyData) *DescribeSystemEventMetaListResponseBody
	GetData() *DescribeSystemEventMetaListResponseBodyData
	SetMessage(v string) *DescribeSystemEventMetaListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSystemEventMetaListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSystemEventMetaListResponseBody
	GetSuccess() *bool
}

type DescribeSystemEventMetaListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried meta information.
	Data *DescribeSystemEventMetaListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6582C8B-E67C-4A19-BC15-EAEFEBDC7995
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSystemEventMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventMetaListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeSystemEventMetaListResponseBody) GetData() *DescribeSystemEventMetaListResponseBodyData {
	return s.Data
}

func (s *DescribeSystemEventMetaListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSystemEventMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemEventMetaListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSystemEventMetaListResponseBody) SetCode(v int32) *DescribeSystemEventMetaListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBody) SetData(v *DescribeSystemEventMetaListResponseBodyData) *DescribeSystemEventMetaListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSystemEventMetaListResponseBody) SetMessage(v string) *DescribeSystemEventMetaListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBody) SetRequestId(v string) *DescribeSystemEventMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBody) SetSuccess(v bool) *DescribeSystemEventMetaListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSystemEventMetaListResponseBodyData struct {
	Resource []*DescribeSystemEventMetaListResponseBodyDataResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventMetaListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventMetaListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventMetaListResponseBodyData) GetResource() []*DescribeSystemEventMetaListResponseBodyDataResource {
	return s.Resource
}

func (s *DescribeSystemEventMetaListResponseBodyData) SetResource(v []*DescribeSystemEventMetaListResponseBodyDataResource) *DescribeSystemEventMetaListResponseBodyData {
	s.Resource = v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSystemEventMetaListResponseBodyDataResource struct {
	// The type of the system event. Valid values:
	//
	// 	- StatusNotification: fault notifications
	//
	// 	- Exception: exceptions
	//
	// 	- Maintenance: O\\&M
	//
	// example:
	//
	// Exception
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The alert level. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
	//
	// example:
	//
	// INFO
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the system event.
	//
	// example:
	//
	// SelectFailureRate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the event name.
	//
	// example:
	//
	// High query failure rate
	NameDesc    *string `json:"NameDesc,omitempty" xml:"NameDesc,omitempty"`
	NameDesc_en *string `json:"NameDesc.En,omitempty" xml:"NameDesc.En,omitempty"`
	// The abbreviation of the service name.
	//
	// example:
	//
	// ADS
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The status of the system event.
	//
	// example:
	//
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the event status.
	//
	// example:
	//
	// Operation Failed
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s DescribeSystemEventMetaListResponseBodyDataResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventMetaListResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetNameDesc() *string {
	return s.NameDesc
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetNameDesc_en() *string {
	return s.NameDesc_en
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetEventType(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetLevel(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetName(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetNameDesc(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.NameDesc = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetNameDesc_en(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.NameDesc_en = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetProduct(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetStatus(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) SetStatusDesc(v string) *DescribeSystemEventMetaListResponseBodyDataResource {
	s.StatusDesc = &v
	return s
}

func (s *DescribeSystemEventMetaListResponseBodyDataResource) Validate() error {
	return dara.Validate(s)
}
