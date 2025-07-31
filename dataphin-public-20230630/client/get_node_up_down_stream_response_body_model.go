// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeUpDownStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNodeUpDownStreamResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetNodeUpDownStreamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetNodeUpDownStreamResponseBody
	GetMessage() *string
	SetNodeDagInfo(v *GetNodeUpDownStreamResponseBodyNodeDagInfo) *GetNodeUpDownStreamResponseBody
	GetNodeDagInfo() *GetNodeUpDownStreamResponseBodyNodeDagInfo
	SetRequestId(v string) *GetNodeUpDownStreamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeUpDownStreamResponseBody
	GetSuccess() *bool
}

type GetNodeUpDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeDagInfo *GetNodeUpDownStreamResponseBodyNodeDagInfo `json:"NodeDagInfo,omitempty" xml:"NodeDagInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeUpDownStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNodeUpDownStreamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeUpDownStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNodeUpDownStreamResponseBody) GetNodeDagInfo() *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	return s.NodeDagInfo
}

func (s *GetNodeUpDownStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeUpDownStreamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeUpDownStreamResponseBody) SetCode(v string) *GetNodeUpDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetHttpStatusCode(v int32) *GetNodeUpDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetMessage(v string) *GetNodeUpDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetNodeDagInfo(v *GetNodeUpDownStreamResponseBodyNodeDagInfo) *GetNodeUpDownStreamResponseBody {
	s.NodeDagInfo = v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetRequestId(v string) *GetNodeUpDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetSuccess(v bool) *GetNodeUpDownStreamResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeUpDownStreamResponseBodyNodeDagInfo struct {
	DownStreamNodeList []*GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList `json:"DownStreamNodeList,omitempty" xml:"DownStreamNodeList,omitempty" type:"Repeated"`
	StartNodeList      []*GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList      `json:"StartNodeList,omitempty" xml:"StartNodeList,omitempty" type:"Repeated"`
	UpStreamNodeList   []*GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList   `json:"UpStreamNodeList,omitempty" xml:"UpStreamNodeList,omitempty" type:"Repeated"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfo) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) GetDownStreamNodeList() []*GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	return s.DownStreamNodeList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) GetStartNodeList() []*GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	return s.StartNodeList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) GetUpStreamNodeList() []*GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	return s.UpStreamNodeList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetDownStreamNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.DownStreamNodeList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetStartNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.StartNodeList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetUpStreamNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.UpStreamNodeList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) Validate() error {
	return dara.Validate(s)
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GetId() *string {
	return s.Id
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GetName() *string {
	return s.Name
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GetType() *string {
	return s.Type
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Type = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) Validate() error {
	return dara.Validate(s)
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GetId() *string {
	return s.Id
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GetName() *string {
	return s.Name
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GetType() *string {
	return s.Type
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Type = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) Validate() error {
	return dara.Validate(s)
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GetId() *string {
	return s.Id
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GetName() *string {
	return s.Name
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GetType() *string {
	return s.Type
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Type = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) Validate() error {
	return dara.Validate(s)
}
