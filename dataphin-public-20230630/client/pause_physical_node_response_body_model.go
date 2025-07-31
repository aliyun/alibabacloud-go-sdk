// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePhysicalNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PausePhysicalNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PausePhysicalNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PausePhysicalNodeResponseBody
	GetMessage() *string
	SetNodeOperateResultList(v []*PausePhysicalNodeResponseBodyNodeOperateResultList) *PausePhysicalNodeResponseBody
	GetNodeOperateResultList() []*PausePhysicalNodeResponseBodyNodeOperateResultList
	SetRequestId(v string) *PausePhysicalNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PausePhysicalNodeResponseBody
	GetSuccess() *bool
}

type PausePhysicalNodeResponseBody struct {
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
	Message               *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeOperateResultList []*PausePhysicalNodeResponseBodyNodeOperateResultList `json:"NodeOperateResultList,omitempty" xml:"NodeOperateResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PausePhysicalNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *PausePhysicalNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PausePhysicalNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PausePhysicalNodeResponseBody) GetNodeOperateResultList() []*PausePhysicalNodeResponseBodyNodeOperateResultList {
	return s.NodeOperateResultList
}

func (s *PausePhysicalNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PausePhysicalNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PausePhysicalNodeResponseBody) SetCode(v string) *PausePhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetHttpStatusCode(v int32) *PausePhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetMessage(v string) *PausePhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetNodeOperateResultList(v []*PausePhysicalNodeResponseBodyNodeOperateResultList) *PausePhysicalNodeResponseBody {
	s.NodeOperateResultList = v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetRequestId(v string) *PausePhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetSuccess(v bool) *PausePhysicalNodeResponseBody {
	s.Success = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type PausePhysicalNodeResponseBodyNodeOperateResultList struct {
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PausePhysicalNodeResponseBodyNodeOperateResultList) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeResponseBodyNodeOperateResultList) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) GetNodeId() *string {
	return s.NodeId
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) GetStatus() *string {
	return s.Status
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetErrorMessage(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.ErrorMessage = &v
	return s
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetNodeId(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.NodeId = &v
	return s
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetStatus(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.Status = &v
	return s
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) Validate() error {
	return dara.Validate(s)
}
