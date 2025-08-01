// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CloneNacosConfigResponseBody
	GetCode() *int32
	SetData(v *CloneNacosConfigResponseBodyData) *CloneNacosConfigResponseBody
	GetData() *CloneNacosConfigResponseBodyData
	SetDynamicMessage(v string) *CloneNacosConfigResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *CloneNacosConfigResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *CloneNacosConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CloneNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloneNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloneNacosConfigResponseBody
	GetSuccess() *bool
}

type CloneNacosConfigResponseBody struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data overview.
	Data *CloneNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Dynamic error message, used to replace the **%s*	- in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it indicates that the input request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// Clone Completed Successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 6678DBA9-5600-5948-ACF8-ED3105B288A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, with the following values:
	//
	// - `true`: Request succeeded.
	//
	// - `false`: Request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloneNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CloneNacosConfigResponseBody) GetData() *CloneNacosConfigResponseBodyData {
	return s.Data
}

func (s *CloneNacosConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CloneNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CloneNacosConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CloneNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloneNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloneNacosConfigResponseBody) SetCode(v int32) *CloneNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetData(v *CloneNacosConfigResponseBodyData) *CloneNacosConfigResponseBody {
	s.Data = v
	return s
}

func (s *CloneNacosConfigResponseBody) SetDynamicMessage(v string) *CloneNacosConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetErrorCode(v string) *CloneNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetHttpStatusCode(v int32) *CloneNacosConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetMessage(v string) *CloneNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetRequestId(v string) *CloneNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetSuccess(v bool) *CloneNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CloneNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type CloneNacosConfigResponseBodyData struct {
	// Failed configuration information.
	FailData []*CloneNacosConfigResponseBodyDataFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Repeated"`
	// Number of skips.
	//
	// example:
	//
	// 1
	SkipCount *int32 `json:"SkipCount,omitempty" xml:"SkipCount,omitempty"`
	// Skipped configuration information.
	SkipData []*CloneNacosConfigResponseBodyDataSkipData `json:"SkipData,omitempty" xml:"SkipData,omitempty" type:"Repeated"`
	// Number of successes.
	//
	// example:
	//
	// 100
	SuccCount *int32 `json:"SuccCount,omitempty" xml:"SuccCount,omitempty"`
}

func (s CloneNacosConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyData) GetFailData() []*CloneNacosConfigResponseBodyDataFailData {
	return s.FailData
}

func (s *CloneNacosConfigResponseBodyData) GetSkipCount() *int32 {
	return s.SkipCount
}

func (s *CloneNacosConfigResponseBodyData) GetSkipData() []*CloneNacosConfigResponseBodyDataSkipData {
	return s.SkipData
}

func (s *CloneNacosConfigResponseBodyData) GetSuccCount() *int32 {
	return s.SuccCount
}

func (s *CloneNacosConfigResponseBodyData) SetFailData(v []*CloneNacosConfigResponseBodyDataFailData) *CloneNacosConfigResponseBodyData {
	s.FailData = v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSkipCount(v int32) *CloneNacosConfigResponseBodyData {
	s.SkipCount = &v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSkipData(v []*CloneNacosConfigResponseBodyDataSkipData) *CloneNacosConfigResponseBodyData {
	s.SkipData = v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSuccCount(v int32) *CloneNacosConfigResponseBodyData {
	s.SuccCount = &v
	return s
}

func (s *CloneNacosConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CloneNacosConfigResponseBodyDataFailData struct {
	// Data ID.
	//
	// example:
	//
	// test2.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Group ID.
	//
	// example:
	//
	// test
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The reason for the current operation.
	//
	// example:
	//
	// param not support
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CloneNacosConfigResponseBodyDataFailData) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigResponseBodyDataFailData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyDataFailData) GetDataId() *string {
	return s.DataId
}

func (s *CloneNacosConfigResponseBodyDataFailData) GetGroup() *string {
	return s.Group
}

func (s *CloneNacosConfigResponseBodyDataFailData) GetReason() *string {
	return s.Reason
}

func (s *CloneNacosConfigResponseBodyDataFailData) SetDataId(v string) *CloneNacosConfigResponseBodyDataFailData {
	s.DataId = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataFailData) SetGroup(v string) *CloneNacosConfigResponseBodyDataFailData {
	s.Group = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataFailData) SetReason(v string) *CloneNacosConfigResponseBodyDataFailData {
	s.Reason = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataFailData) Validate() error {
	return dara.Validate(s)
}

type CloneNacosConfigResponseBodyDataSkipData struct {
	// Data ID.
	//
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Group ID.
	//
	// example:
	//
	// public
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s CloneNacosConfigResponseBodyDataSkipData) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigResponseBodyDataSkipData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyDataSkipData) GetDataId() *string {
	return s.DataId
}

func (s *CloneNacosConfigResponseBodyDataSkipData) GetGroup() *string {
	return s.Group
}

func (s *CloneNacosConfigResponseBodyDataSkipData) SetDataId(v string) *CloneNacosConfigResponseBodyDataSkipData {
	s.DataId = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataSkipData) SetGroup(v string) *CloneNacosConfigResponseBodyDataSkipData {
	s.Group = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataSkipData) Validate() error {
	return dara.Validate(s)
}
