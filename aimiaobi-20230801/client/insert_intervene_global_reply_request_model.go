// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneGlobalReplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *InsertInterveneGlobalReplyRequest
	GetAgentKey() *string
	SetReplyMessagList(v []*InsertInterveneGlobalReplyRequestReplyMessagList) *InsertInterveneGlobalReplyRequest
	GetReplyMessagList() []*InsertInterveneGlobalReplyRequestReplyMessagList
}

type InsertInterveneGlobalReplyRequest struct {
	// Unique identifier of the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// List of reply content
	ReplyMessagList []*InsertInterveneGlobalReplyRequestReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s InsertInterveneGlobalReplyRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *InsertInterveneGlobalReplyRequest) GetReplyMessagList() []*InsertInterveneGlobalReplyRequestReplyMessagList {
	return s.ReplyMessagList
}

func (s *InsertInterveneGlobalReplyRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequest) SetReplyMessagList(v []*InsertInterveneGlobalReplyRequestReplyMessagList) *InsertInterveneGlobalReplyRequest {
	s.ReplyMessagList = v
	return s
}

func (s *InsertInterveneGlobalReplyRequest) Validate() error {
	if s.ReplyMessagList != nil {
		for _, item := range s.ReplyMessagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsertInterveneGlobalReplyRequestReplyMessagList struct {
	// Reply content
	//
	// example:
	//
	// 抱歉，问题我无法回答
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Reply type
	//
	// example:
	//
	// 全局回复的类型
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) GetMessage() *string {
	return s.Message
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) GetReplyType() *string {
	return s.ReplyType
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetMessage(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetReplyType(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.ReplyType = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) Validate() error {
	return dara.Validate(s)
}
