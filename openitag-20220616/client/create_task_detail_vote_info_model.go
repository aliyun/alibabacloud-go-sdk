// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskDetailVoteInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMinVote(v int64) *CreateTaskDetailVoteInfo
	GetMinVote() *int64
	SetVoteNum(v int64) *CreateTaskDetailVoteInfo
	GetVoteNum() *int64
}

type CreateTaskDetailVoteInfo struct {
	// example:
	//
	// 3
	MinVote *int64 `json:"MinVote,omitempty" xml:"MinVote,omitempty"`
	// example:
	//
	// 3
	VoteNum *int64 `json:"VoteNum,omitempty" xml:"VoteNum,omitempty"`
}

func (s CreateTaskDetailVoteInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskDetailVoteInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailVoteInfo) GetMinVote() *int64 {
	return s.MinVote
}

func (s *CreateTaskDetailVoteInfo) GetVoteNum() *int64 {
	return s.VoteNum
}

func (s *CreateTaskDetailVoteInfo) SetMinVote(v int64) *CreateTaskDetailVoteInfo {
	s.MinVote = &v
	return s
}

func (s *CreateTaskDetailVoteInfo) SetVoteNum(v int64) *CreateTaskDetailVoteInfo {
	s.VoteNum = &v
	return s
}

func (s *CreateTaskDetailVoteInfo) Validate() error {
	return dara.Validate(s)
}
