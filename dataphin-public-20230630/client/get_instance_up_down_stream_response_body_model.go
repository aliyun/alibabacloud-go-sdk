// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUpDownStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceUpDownStreamResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetInstanceUpDownStreamResponseBody
	GetHttpStatusCode() *int32
	SetInstanceDagInfo(v *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) *GetInstanceUpDownStreamResponseBody
	GetInstanceDagInfo() *GetInstanceUpDownStreamResponseBodyInstanceDagInfo
	SetMessage(v string) *GetInstanceUpDownStreamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceUpDownStreamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceUpDownStreamResponseBody
	GetSuccess() *bool
}

type GetInstanceUpDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode  *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceDagInfo *GetInstanceUpDownStreamResponseBodyInstanceDagInfo `json:"InstanceDagInfo,omitempty" xml:"InstanceDagInfo,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceUpDownStreamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceUpDownStreamResponseBody) GetInstanceDagInfo() *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	return s.InstanceDagInfo
}

func (s *GetInstanceUpDownStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceUpDownStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceUpDownStreamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceUpDownStreamResponseBody) SetCode(v string) *GetInstanceUpDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetHttpStatusCode(v int32) *GetInstanceUpDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetInstanceDagInfo(v *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) *GetInstanceUpDownStreamResponseBody {
	s.InstanceDagInfo = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetMessage(v string) *GetInstanceUpDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetRequestId(v string) *GetInstanceUpDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetSuccess(v bool) *GetInstanceUpDownStreamResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfo struct {
	DownInstanceList  []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList  `json:"DownInstanceList,omitempty" xml:"DownInstanceList,omitempty" type:"Repeated"`
	StartInstanceList []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList `json:"StartInstanceList,omitempty" xml:"StartInstanceList,omitempty" type:"Repeated"`
	UpInstanceList    []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList    `json:"UpInstanceList,omitempty" xml:"UpInstanceList,omitempty" type:"Repeated"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) GetDownInstanceList() []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	return s.DownInstanceList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) GetStartInstanceList() []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	return s.StartInstanceList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) GetUpInstanceList() []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	return s.UpInstanceList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetDownInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.DownInstanceList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetStartInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.StartInstanceList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetUpInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.UpInstanceList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GetId() *string {
	return s.Id
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GetName() *string {
	return s.Name
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.NodeType = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GetId() *string {
	return s.Id
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GetName() *string {
	return s.Name
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.NodeType = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GetId() *string {
	return s.Id
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GetName() *string {
	return s.Name
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.NodeType = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) Validate() error {
	return dara.Validate(s)
}
