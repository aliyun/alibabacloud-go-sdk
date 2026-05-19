// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectBindMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*string) *ListFileProtectBindMachineResponseBody
	GetList() []*string
	SetPageInfo(v *ListFileProtectBindMachineResponseBodyPageInfo) *ListFileProtectBindMachineResponseBody
	GetPageInfo() *ListFileProtectBindMachineResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectBindMachineResponseBody
	GetRequestId() *string
}

type ListFileProtectBindMachineResponseBody struct {
	List     []*string                                       `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *ListFileProtectBindMachineResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectBindMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectBindMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectBindMachineResponseBody) GetList() []*string {
	return s.List
}

func (s *ListFileProtectBindMachineResponseBody) GetPageInfo() *ListFileProtectBindMachineResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectBindMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectBindMachineResponseBody) SetList(v []*string) *ListFileProtectBindMachineResponseBody {
	s.List = v
	return s
}

func (s *ListFileProtectBindMachineResponseBody) SetPageInfo(v *ListFileProtectBindMachineResponseBodyPageInfo) *ListFileProtectBindMachineResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectBindMachineResponseBody) SetRequestId(v string) *ListFileProtectBindMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectBindMachineResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFileProtectBindMachineResponseBodyPageInfo struct {
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 69
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectBindMachineResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectBindMachineResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectBindMachineResponseBodyPageInfo) GetCount() *int64 {
	return s.Count
}

func (s *ListFileProtectBindMachineResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFileProtectBindMachineResponseBodyPageInfo) SetCount(v int64) *ListFileProtectBindMachineResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListFileProtectBindMachineResponseBodyPageInfo) SetTotalCount(v int64) *ListFileProtectBindMachineResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectBindMachineResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
