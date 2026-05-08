// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CountTextResponseBody
	GetRequestId() *string
	SetCountTextCmdList(v []*CountTextResponseBodyCountTextCmdList) *CountTextResponseBody
	GetCountTextCmdList() []*CountTextResponseBodyCountTextCmdList
}

type CountTextResponseBody struct {
	// example:
	//
	// 6C9CB64D-E2D3-5BF2-A9E6-2445F952F178
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CountTextCmdList []*CountTextResponseBodyCountTextCmdList `json:"countTextCmdList,omitempty" xml:"countTextCmdList,omitempty" type:"Repeated"`
}

func (s CountTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountTextResponseBody) GoString() string {
	return s.String()
}

func (s *CountTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountTextResponseBody) GetCountTextCmdList() []*CountTextResponseBodyCountTextCmdList {
	return s.CountTextCmdList
}

func (s *CountTextResponseBody) SetRequestId(v string) *CountTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountTextResponseBody) SetCountTextCmdList(v []*CountTextResponseBodyCountTextCmdList) *CountTextResponseBody {
	s.CountTextCmdList = v
	return s
}

func (s *CountTextResponseBody) Validate() error {
	if s.CountTextCmdList != nil {
		for _, item := range s.CountTextCmdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CountTextResponseBodyCountTextCmdList struct {
	// example:
	//
	// 4
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// RED_BOOK
	Theme *string `json:"theme,omitempty" xml:"theme,omitempty"`
}

func (s CountTextResponseBodyCountTextCmdList) String() string {
	return dara.Prettify(s)
}

func (s CountTextResponseBodyCountTextCmdList) GoString() string {
	return s.String()
}

func (s *CountTextResponseBodyCountTextCmdList) GetCount() *int64 {
	return s.Count
}

func (s *CountTextResponseBodyCountTextCmdList) GetTheme() *string {
	return s.Theme
}

func (s *CountTextResponseBodyCountTextCmdList) SetCount(v int64) *CountTextResponseBodyCountTextCmdList {
	s.Count = &v
	return s
}

func (s *CountTextResponseBodyCountTextCmdList) SetTheme(v string) *CountTextResponseBodyCountTextCmdList {
	s.Theme = &v
	return s
}

func (s *CountTextResponseBodyCountTextCmdList) Validate() error {
	return dara.Validate(s)
}
