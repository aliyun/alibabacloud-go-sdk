// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventNameList(v []*DescribeInvadeEventNameListResponseBodyEventNameList) *DescribeInvadeEventNameListResponseBody
	GetEventNameList() []*DescribeInvadeEventNameListResponseBodyEventNameList
	SetNameList(v []*string) *DescribeInvadeEventNameListResponseBody
	GetNameList() []*string
	SetRequestId(v string) *DescribeInvadeEventNameListResponseBody
	GetRequestId() *string
}

type DescribeInvadeEventNameListResponseBody struct {
	EventNameList []*DescribeInvadeEventNameListResponseBodyEventNameList `json:"EventNameList,omitempty" xml:"EventNameList,omitempty" type:"Repeated"`
	NameList      []*string                                               `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	// example:
	//
	// 6ABAA264-E7B5-5D66-8FC3-9253100****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvadeEventNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventNameListResponseBody) GetEventNameList() []*DescribeInvadeEventNameListResponseBodyEventNameList {
	return s.EventNameList
}

func (s *DescribeInvadeEventNameListResponseBody) GetNameList() []*string {
	return s.NameList
}

func (s *DescribeInvadeEventNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEventNameListResponseBody) SetEventNameList(v []*DescribeInvadeEventNameListResponseBodyEventNameList) *DescribeInvadeEventNameListResponseBody {
	s.EventNameList = v
	return s
}

func (s *DescribeInvadeEventNameListResponseBody) SetNameList(v []*string) *DescribeInvadeEventNameListResponseBody {
	s.NameList = v
	return s
}

func (s *DescribeInvadeEventNameListResponseBody) SetRequestId(v string) *DescribeInvadeEventNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEventNameListResponseBody) Validate() error {
	if s.EventNameList != nil {
		for _, item := range s.EventNameList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvadeEventNameListResponseBodyEventNameList struct {
	// example:
	//
	// CActivity
	EventKey  *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
}

func (s DescribeInvadeEventNameListResponseBodyEventNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventNameListResponseBodyEventNameList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventNameListResponseBodyEventNameList) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeInvadeEventNameListResponseBodyEventNameList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeInvadeEventNameListResponseBodyEventNameList) SetEventKey(v string) *DescribeInvadeEventNameListResponseBodyEventNameList {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventNameListResponseBodyEventNameList) SetEventName(v string) *DescribeInvadeEventNameListResponseBodyEventNameList {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventNameListResponseBodyEventNameList) Validate() error {
	return dara.Validate(s)
}
