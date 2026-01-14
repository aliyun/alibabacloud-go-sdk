// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *GetMailTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetMailTaskListRequest
	GetPageSize() *int32
	SetPaused(v bool) *GetMailTaskListRequest
	GetPaused() *bool
	SetUserNick(v string) *GetMailTaskListRequest
	GetUserNick() *string
}

type GetMailTaskListRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// example:
	//
	// test
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s GetMailTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetMailTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMailTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMailTaskListRequest) GetPaused() *bool {
	return s.Paused
}

func (s *GetMailTaskListRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *GetMailTaskListRequest) SetPageNum(v int32) *GetMailTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *GetMailTaskListRequest) SetPageSize(v int32) *GetMailTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetMailTaskListRequest) SetPaused(v bool) *GetMailTaskListRequest {
	s.Paused = &v
	return s
}

func (s *GetMailTaskListRequest) SetUserNick(v string) *GetMailTaskListRequest {
	s.UserNick = &v
	return s
}

func (s *GetMailTaskListRequest) Validate() error {
	return dara.Validate(s)
}
