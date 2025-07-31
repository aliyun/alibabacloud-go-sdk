// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDownStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNodeDownStreamResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListNodeDownStreamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNodeDownStreamResponseBody
	GetMessage() *string
	SetNodeInfoList(v []*ListNodeDownStreamResponseBodyNodeInfoList) *ListNodeDownStreamResponseBody
	GetNodeInfoList() []*ListNodeDownStreamResponseBodyNodeInfoList
	SetRequestId(v string) *ListNodeDownStreamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodeDownStreamResponseBody
	GetSuccess() *bool
}

type ListNodeDownStreamResponseBody struct {
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
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfoList []*ListNodeDownStreamResponseBodyNodeInfoList `json:"NodeInfoList,omitempty" xml:"NodeInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeDownStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNodeDownStreamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodeDownStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNodeDownStreamResponseBody) GetNodeInfoList() []*ListNodeDownStreamResponseBodyNodeInfoList {
	return s.NodeInfoList
}

func (s *ListNodeDownStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeDownStreamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodeDownStreamResponseBody) SetCode(v string) *ListNodeDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetHttpStatusCode(v int32) *ListNodeDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetMessage(v string) *ListNodeDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetNodeInfoList(v []*ListNodeDownStreamResponseBodyNodeInfoList) *ListNodeDownStreamResponseBody {
	s.NodeInfoList = v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetRequestId(v string) *ListNodeDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetSuccess(v bool) *ListNodeDownStreamResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodeDownStreamResponseBodyNodeInfoList struct {
	// example:
	//
	// 1
	Depth       *int32    `json:"Depth,omitempty" xml:"Depth,omitempty"`
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_2423351
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDownStreamResponseBodyNodeInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamResponseBodyNodeInfoList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) GetDepth() *int32 {
	return s.Depth
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) GetFieldIdList() []*string {
	return s.FieldIdList
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) GetId() *string {
	return s.Id
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) GetName() *string {
	return s.Name
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) GetType() *string {
	return s.Type
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetDepth(v int32) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Depth = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetFieldIdList(v []*string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.FieldIdList = v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetId(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Id = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetName(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Name = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetType(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Type = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) Validate() error {
	return dara.Validate(s)
}
