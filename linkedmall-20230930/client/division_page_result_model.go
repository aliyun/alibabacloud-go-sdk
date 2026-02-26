// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDivisionPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionList(v []*Division) *DivisionPageResult
	GetDivisionList() []*Division
	SetRequestId(v string) *DivisionPageResult
	GetRequestId() *string
}

type DivisionPageResult struct {
	DivisionList []*Division `json:"divisionList,omitempty" xml:"divisionList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DivisionPageResult) String() string {
	return dara.Prettify(s)
}

func (s DivisionPageResult) GoString() string {
	return s.String()
}

func (s *DivisionPageResult) GetDivisionList() []*Division {
	return s.DivisionList
}

func (s *DivisionPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DivisionPageResult) SetDivisionList(v []*Division) *DivisionPageResult {
	s.DivisionList = v
	return s
}

func (s *DivisionPageResult) SetRequestId(v string) *DivisionPageResult {
	s.RequestId = &v
	return s
}

func (s *DivisionPageResult) Validate() error {
	if s.DivisionList != nil {
		for _, item := range s.DivisionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
