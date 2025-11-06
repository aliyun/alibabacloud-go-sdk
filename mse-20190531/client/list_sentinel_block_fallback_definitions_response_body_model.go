// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSentinelBlockFallbackDefinitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSentinelBlockFallbackDefinitionsResponseBodyData) *ListSentinelBlockFallbackDefinitionsResponseBody
	GetData() []*ListSentinelBlockFallbackDefinitionsResponseBodyData
	SetHttpStatusCode(v int32) *ListSentinelBlockFallbackDefinitionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSentinelBlockFallbackDefinitionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSentinelBlockFallbackDefinitionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSentinelBlockFallbackDefinitionsResponseBody
	GetSuccess() *bool
}

type ListSentinelBlockFallbackDefinitionsResponseBody struct {
	// The details of the data.
	Data []*ListSentinelBlockFallbackDefinitionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSentinelBlockFallbackDefinitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSentinelBlockFallbackDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) GetData() []*ListSentinelBlockFallbackDefinitionsResponseBodyData {
	return s.Data
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) SetData(v []*ListSentinelBlockFallbackDefinitionsResponseBodyData) *ListSentinelBlockFallbackDefinitionsResponseBody {
	s.Data = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) SetHttpStatusCode(v int32) *ListSentinelBlockFallbackDefinitionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) SetMessage(v string) *ListSentinelBlockFallbackDefinitionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) SetRequestId(v string) *ListSentinelBlockFallbackDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) SetSuccess(v bool) *ListSentinelBlockFallbackDefinitionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSentinelBlockFallbackDefinitionsResponseBodyData struct {
	// The name of the application.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Behavior  detail.
	//
	// example:
	//
	// {"webRespStatusCode":429,"webRespMessage":"test","webFallbackMode":0,"webRespContentType":0}
	FallbackBehavior map[string]interface{} `json:"FallbackBehavior,omitempty" xml:"FallbackBehavior,omitempty"`
	// Behavior Id
	//
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the behavior.
	//
	// example:
	//
	// defaultFallback
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Microservices namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Behavior classification.
	//
	// example:
	//
	// 1
	ResourceClassification *string `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// Resource information bound to the behavior.
	//
	// example:
	//
	// {"/params/{hot}":[1]}
	TargetMap map[string]interface{} `json:"TargetMap,omitempty" xml:"TargetMap,omitempty"`
}

func (s ListSentinelBlockFallbackDefinitionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSentinelBlockFallbackDefinitionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetFallbackBehavior() map[string]interface{} {
	return s.FallbackBehavior
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetResourceClassification() *string {
	return s.ResourceClassification
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) GetTargetMap() map[string]interface{} {
	return s.TargetMap
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetAppName(v string) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetFallbackBehavior(v map[string]interface{}) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.FallbackBehavior = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetId(v string) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetName(v string) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetNamespace(v string) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetResourceClassification(v string) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.ResourceClassification = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) SetTargetMap(v map[string]interface{}) *ListSentinelBlockFallbackDefinitionsResponseBodyData {
	s.TargetMap = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
