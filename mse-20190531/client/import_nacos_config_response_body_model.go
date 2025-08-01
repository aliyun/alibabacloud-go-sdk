// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportNacosConfigResponseBody
	GetCode() *int32
	SetData(v *ImportNacosConfigResponseBodyData) *ImportNacosConfigResponseBody
	GetData() *ImportNacosConfigResponseBodyData
	SetDynamicMessage(v string) *ImportNacosConfigResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ImportNacosConfigResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ImportNacosConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportNacosConfigResponseBody
	GetSuccess() *bool
}

type ImportNacosConfigResponseBody struct {
	// The error message returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of configurations that are imported.
	Data *ImportNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The details of the data.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The code returned.
	//
	// example:
	//
	// mse-100-100
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request is successfully processed.
	//
	// example:
	//
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// AF21683A-29C7-4853-AC0F-B5ADEE4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportNacosConfigResponseBody) GetData() *ImportNacosConfigResponseBodyData {
	return s.Data
}

func (s *ImportNacosConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ImportNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ImportNacosConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportNacosConfigResponseBody) SetCode(v int32) *ImportNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetData(v *ImportNacosConfigResponseBodyData) *ImportNacosConfigResponseBody {
	s.Data = v
	return s
}

func (s *ImportNacosConfigResponseBody) SetDynamicMessage(v string) *ImportNacosConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetErrorCode(v string) *ImportNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetHttpStatusCode(v int32) *ImportNacosConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetMessage(v string) *ImportNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetRequestId(v string) *ImportNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetSuccess(v bool) *ImportNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ImportNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportNacosConfigResponseBodyData struct {
	// The data structure.
	FailData []*ImportNacosConfigResponseBodyDataFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Repeated"`
	// The information about skipped configurations.
	//
	// example:
	//
	// 10
	SkipCount *int32 `json:"SkipCount,omitempty" xml:"SkipCount,omitempty"`
	// The data structure.
	SkipData []*ImportNacosConfigResponseBodyDataSkipData `json:"SkipData,omitempty" xml:"SkipData,omitempty" type:"Repeated"`
	// The number of configurations that are skipped.
	//
	// example:
	//
	// 100
	SuccCount *int32 `json:"SuccCount,omitempty" xml:"SuccCount,omitempty"`
}

func (s ImportNacosConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyData) GetFailData() []*ImportNacosConfigResponseBodyDataFailData {
	return s.FailData
}

func (s *ImportNacosConfigResponseBodyData) GetSkipCount() *int32 {
	return s.SkipCount
}

func (s *ImportNacosConfigResponseBodyData) GetSkipData() []*ImportNacosConfigResponseBodyDataSkipData {
	return s.SkipData
}

func (s *ImportNacosConfigResponseBodyData) GetSuccCount() *int32 {
	return s.SuccCount
}

func (s *ImportNacosConfigResponseBodyData) SetFailData(v []*ImportNacosConfigResponseBodyDataFailData) *ImportNacosConfigResponseBodyData {
	s.FailData = v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSkipCount(v int32) *ImportNacosConfigResponseBodyData {
	s.SkipCount = &v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSkipData(v []*ImportNacosConfigResponseBodyDataSkipData) *ImportNacosConfigResponseBodyData {
	s.SkipData = v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSuccCount(v int32) *ImportNacosConfigResponseBodyData {
	s.SuccCount = &v
	return s
}

func (s *ImportNacosConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ImportNacosConfigResponseBodyDataFailData struct {
	// The ID of the group.
	//
	// example:
	//
	// test2.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// test
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ImportNacosConfigResponseBodyDataFailData) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigResponseBodyDataFailData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyDataFailData) GetDataId() *string {
	return s.DataId
}

func (s *ImportNacosConfigResponseBodyDataFailData) GetGroup() *string {
	return s.Group
}

func (s *ImportNacosConfigResponseBodyDataFailData) GetReason() *string {
	return s.Reason
}

func (s *ImportNacosConfigResponseBodyDataFailData) SetDataId(v string) *ImportNacosConfigResponseBodyDataFailData {
	s.DataId = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataFailData) SetGroup(v string) *ImportNacosConfigResponseBodyDataFailData {
	s.Group = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataFailData) SetReason(v string) *ImportNacosConfigResponseBodyDataFailData {
	s.Reason = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataFailData) Validate() error {
	return dara.Validate(s)
}

type ImportNacosConfigResponseBodyDataSkipData struct {
	// The ID of the group.
	//
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The information about configurations that are failed to be imported.
	//
	// example:
	//
	// public
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s ImportNacosConfigResponseBodyDataSkipData) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigResponseBodyDataSkipData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyDataSkipData) GetDataId() *string {
	return s.DataId
}

func (s *ImportNacosConfigResponseBodyDataSkipData) GetGroup() *string {
	return s.Group
}

func (s *ImportNacosConfigResponseBodyDataSkipData) SetDataId(v string) *ImportNacosConfigResponseBodyDataSkipData {
	s.DataId = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataSkipData) SetGroup(v string) *ImportNacosConfigResponseBodyDataSkipData {
	s.Group = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataSkipData) Validate() error {
	return dara.Validate(s)
}
