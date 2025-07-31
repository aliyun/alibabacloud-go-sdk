// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePhysicalNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumePhysicalNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ResumePhysicalNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResumePhysicalNodeResponseBody
	GetMessage() *string
	SetNodeOperateResultList(v []*ResumePhysicalNodeResponseBodyNodeOperateResultList) *ResumePhysicalNodeResponseBody
	GetNodeOperateResultList() []*ResumePhysicalNodeResponseBodyNodeOperateResultList
	SetRequestId(v string) *ResumePhysicalNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumePhysicalNodeResponseBody
	GetSuccess() *bool
}

type ResumePhysicalNodeResponseBody struct {
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
	Message               *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeOperateResultList []*ResumePhysicalNodeResponseBodyNodeOperateResultList `json:"NodeOperateResultList,omitempty" xml:"NodeOperateResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumePhysicalNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumePhysicalNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResumePhysicalNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumePhysicalNodeResponseBody) GetNodeOperateResultList() []*ResumePhysicalNodeResponseBodyNodeOperateResultList {
	return s.NodeOperateResultList
}

func (s *ResumePhysicalNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumePhysicalNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumePhysicalNodeResponseBody) SetCode(v string) *ResumePhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetHttpStatusCode(v int32) *ResumePhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetMessage(v string) *ResumePhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetNodeOperateResultList(v []*ResumePhysicalNodeResponseBodyNodeOperateResultList) *ResumePhysicalNodeResponseBody {
	s.NodeOperateResultList = v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetRequestId(v string) *ResumePhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetSuccess(v bool) *ResumePhysicalNodeResponseBody {
	s.Success = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ResumePhysicalNodeResponseBodyNodeOperateResultList struct {
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

func (s ResumePhysicalNodeResponseBodyNodeOperateResultList) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeResponseBodyNodeOperateResultList) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) GetNodeId() *string {
	return s.NodeId
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) GetStatus() *string {
	return s.Status
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetErrorMessage(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.ErrorMessage = &v
	return s
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetNodeId(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.NodeId = &v
	return s
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetStatus(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.Status = &v
	return s
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) Validate() error {
	return dara.Validate(s)
}
