// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDialogueNodeStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDialogueNodeStatisticsResponseBody
	GetCode() *string
	SetGroupId(v string) *DescribeDialogueNodeStatisticsResponseBody
	GetGroupId() *string
	SetHangUpDialogueNodes(v []*DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) *DescribeDialogueNodeStatisticsResponseBody
	GetHangUpDialogueNodes() []*DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes
	SetHttpStatusCode(v int32) *DescribeDialogueNodeStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DescribeDialogueNodeStatisticsResponseBody
	GetInstanceId() *string
	SetMessage(v string) *DescribeDialogueNodeStatisticsResponseBody
	GetMessage() *string
	SetNoAnswerDialogueNodes(v []*DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) *DescribeDialogueNodeStatisticsResponseBody
	GetNoAnswerDialogueNodes() []*DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes
	SetRequestId(v string) *DescribeDialogueNodeStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDialogueNodeStatisticsResponseBody
	GetSuccess() *bool
	SetTotalCompleted(v int32) *DescribeDialogueNodeStatisticsResponseBody
	GetTotalCompleted() *int32
}

type DescribeDialogueNodeStatisticsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// f06f7c9f-2895-4b30-a8c2-6ecccb9c9f89
	GroupId             *string                                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HangUpDialogueNodes []*DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes `json:"HangUpDialogueNodes,omitempty" xml:"HangUpDialogueNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// []
	NoAnswerDialogueNodes []*DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes `json:"NoAnswerDialogueNodes,omitempty" xml:"NoAnswerDialogueNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 13
	TotalCompleted *int32 `json:"TotalCompleted,omitempty" xml:"TotalCompleted,omitempty"`
}

func (s DescribeDialogueNodeStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDialogueNodeStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetHangUpDialogueNodes() []*DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes {
	return s.HangUpDialogueNodes
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetNoAnswerDialogueNodes() []*DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	return s.NoAnswerDialogueNodes
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDialogueNodeStatisticsResponseBody) GetTotalCompleted() *int32 {
	return s.TotalCompleted
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetCode(v string) *DescribeDialogueNodeStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetGroupId(v string) *DescribeDialogueNodeStatisticsResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetHangUpDialogueNodes(v []*DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) *DescribeDialogueNodeStatisticsResponseBody {
	s.HangUpDialogueNodes = v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetHttpStatusCode(v int32) *DescribeDialogueNodeStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetInstanceId(v string) *DescribeDialogueNodeStatisticsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetMessage(v string) *DescribeDialogueNodeStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetNoAnswerDialogueNodes(v []*DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) *DescribeDialogueNodeStatisticsResponseBody {
	s.NoAnswerDialogueNodes = v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetRequestId(v string) *DescribeDialogueNodeStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetSuccess(v bool) *DescribeDialogueNodeStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) SetTotalCompleted(v int32) *DescribeDialogueNodeStatisticsResponseBody {
	s.TotalCompleted = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBody) Validate() error {
	if s.HangUpDialogueNodes != nil {
		for _, item := range s.HangUpDialogueNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NoAnswerDialogueNodes != nil {
		for _, item := range s.NoAnswerDialogueNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes struct {
	HangUpNum   *int32  `json:"HangUpNum,omitempty" xml:"HangUpNum,omitempty"`
	NodeId      *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName    *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RateDisplay *string `json:"RateDisplay,omitempty" xml:"RateDisplay,omitempty"`
}

func (s DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) GoString() string {
	return s.String()
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) GetHangUpNum() *int32 {
	return s.HangUpNum
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) GetRateDisplay() *string {
	return s.RateDisplay
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) SetHangUpNum(v int32) *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes {
	s.HangUpNum = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) SetNodeId(v string) *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) SetNodeName(v string) *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) SetRateDisplay(v string) *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes {
	s.RateDisplay = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyHangUpDialogueNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes struct {
	// example:
	//
	// f06f7c9f-2895-4b30-a8c2-6ecccb9c9f89
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 2
	HangUpNum *int32 `json:"HangUpNum,omitempty" xml:"HangUpNum,omitempty"`
	// example:
	//
	// 12
	HitNum *int32 `json:"HitNum,omitempty" xml:"HitNum,omitempty"`
	// id
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	NoAnswerNum *int32 `json:"NoAnswerNum,omitempty" xml:"NoAnswerNum,omitempty"`
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// xxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GoString() string {
	return s.String()
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetHangUpNum() *int32 {
	return s.HangUpNum
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetHitNum() *int32 {
	return s.HitNum
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetId() *string {
	return s.Id
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetNoAnswerNum() *int32 {
	return s.NoAnswerNum
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetGroupId(v string) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetHangUpNum(v int32) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.HangUpNum = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetHitNum(v int32) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.HitNum = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetId(v string) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.Id = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetInstanceId(v string) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetNoAnswerNum(v int32) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.NoAnswerNum = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetNodeId(v string) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) SetNodeName(v string) *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponseBodyNoAnswerDialogueNodes) Validate() error {
	return dara.Validate(s)
}
