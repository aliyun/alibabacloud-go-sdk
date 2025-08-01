// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSentinelBlockFallbackDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetCode() *int32
	SetData(v *CreateSentinelBlockFallbackDefinitionResponseBodyData) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetData() *CreateSentinelBlockFallbackDefinitionResponseBodyData
	SetHttpStatusCode(v int32) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSentinelBlockFallbackDefinitionResponseBody
	GetSuccess() *string
}

type CreateSentinelBlockFallbackDefinitionResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateSentinelBlockFallbackDefinitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ADDD8AB7-8D1C-4697-A83E-410D2607****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSentinelBlockFallbackDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSentinelBlockFallbackDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetData() *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	return s.Data
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetCode(v int32) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetData(v *CreateSentinelBlockFallbackDefinitionResponseBodyData) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.Data = v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetHttpStatusCode(v int32) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetMessage(v string) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetRequestId(v string) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) SetSuccess(v string) *CreateSentinelBlockFallbackDefinitionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSentinelBlockFallbackDefinitionResponseBodyData struct {
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// {"webRespStatusCode":429,"webRespMessage":"test","webFallbackMode":0,"webRespContentType":0}
	FallbackBehavior *string `json:"FallbackBehavior,omitempty" xml:"FallbackBehavior,omitempty"`
	// example:
	//
	// 34726
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	ResourceClassification *int32 `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateSentinelBlockFallbackDefinitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSentinelBlockFallbackDefinitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetFallbackBehavior() *string {
	return s.FallbackBehavior
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetResourceClassification() *int32 {
	return s.ResourceClassification
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetAppName(v string) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetFallbackBehavior(v string) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.FallbackBehavior = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetId(v int64) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetName(v string) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetNamespace(v string) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetResourceClassification(v int32) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.ResourceClassification = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) SetUserId(v string) *CreateSentinelBlockFallbackDefinitionResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
