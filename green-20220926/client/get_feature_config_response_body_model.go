// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetFeatureConfigResponseBody
	GetCode() *int32
	SetData(v *GetFeatureConfigResponseBodyData) *GetFeatureConfigResponseBody
	GetData() *GetFeatureConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetFeatureConfigResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *GetFeatureConfigResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetFeatureConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFeatureConfigResponseBody
	GetSuccess() *bool
}

type GetFeatureConfigResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *GetFeatureConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message for this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFeatureConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetFeatureConfigResponseBody) GetData() *GetFeatureConfigResponseBodyData {
	return s.Data
}

func (s *GetFeatureConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFeatureConfigResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetFeatureConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFeatureConfigResponseBody) SetCode(v int32) *GetFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetData(v *GetFeatureConfigResponseBodyData) *GetFeatureConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetFeatureConfigResponseBody) SetHttpStatusCode(v int32) *GetFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetMsg(v string) *GetFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetRequestId(v string) *GetFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetSuccess(v bool) *GetFeatureConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetFeatureConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFeatureConfigResponseBodyData struct {
	// List of feature configurations
	FeatureConf []map[string]interface{} `json:"FeatureConf,omitempty" xml:"FeatureConf,omitempty" type:"Repeated"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// UID.
	//
	// example:
	//
	// 1643953****74290
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetFeatureConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponseBodyData) GetFeatureConf() []map[string]interface{} {
	return s.FeatureConf
}

func (s *GetFeatureConfigResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureConfigResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetFeatureConfigResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetFeatureConfigResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetFeatureConfigResponseBodyData) SetFeatureConf(v []map[string]interface{}) *GetFeatureConfigResponseBodyData {
	s.FeatureConf = v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetResourceType(v string) *GetFeatureConfigResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetServiceCode(v string) *GetFeatureConfigResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetType(v string) *GetFeatureConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetUid(v string) *GetFeatureConfigResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
