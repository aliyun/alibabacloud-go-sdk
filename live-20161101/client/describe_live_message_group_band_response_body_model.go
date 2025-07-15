// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupBandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBannedUserList(v []*string) *DescribeLiveMessageGroupBandResponseBody
	GetBannedUserList() []*string
	SetGroupId(v string) *DescribeLiveMessageGroupBandResponseBody
	GetGroupId() *string
	SetIsbannedAll(v bool) *DescribeLiveMessageGroupBandResponseBody
	GetIsbannedAll() *bool
	SetRequestId(v string) *DescribeLiveMessageGroupBandResponseBody
	GetRequestId() *string
	SetUnbannedUserList(v []*string) *DescribeLiveMessageGroupBandResponseBody
	GetUnbannedUserList() []*string
}

type DescribeLiveMessageGroupBandResponseBody struct {
	// The list of users that were muted separately, but not by muting the entire group.
	BannedUserList []*string `json:"BannedUserList,omitempty" xml:"BannedUserList,omitempty" type:"Repeated"`
	// The group ID.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether all users in the interactive messaging group are muted.
	//
	// example:
	//
	// false
	IsbannedAll *bool `json:"IsbannedAll,omitempty" xml:"IsbannedAll,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 021D1FE7-2E87-16AC-9364-4E7EA47C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of users who were not muted when the entire group was muted.
	UnbannedUserList []*string `json:"UnbannedUserList,omitempty" xml:"UnbannedUserList,omitempty" type:"Repeated"`
}

func (s DescribeLiveMessageGroupBandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupBandResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupBandResponseBody) GetBannedUserList() []*string {
	return s.BannedUserList
}

func (s *DescribeLiveMessageGroupBandResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeLiveMessageGroupBandResponseBody) GetIsbannedAll() *bool {
	return s.IsbannedAll
}

func (s *DescribeLiveMessageGroupBandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveMessageGroupBandResponseBody) GetUnbannedUserList() []*string {
	return s.UnbannedUserList
}

func (s *DescribeLiveMessageGroupBandResponseBody) SetBannedUserList(v []*string) *DescribeLiveMessageGroupBandResponseBody {
	s.BannedUserList = v
	return s
}

func (s *DescribeLiveMessageGroupBandResponseBody) SetGroupId(v string) *DescribeLiveMessageGroupBandResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeLiveMessageGroupBandResponseBody) SetIsbannedAll(v bool) *DescribeLiveMessageGroupBandResponseBody {
	s.IsbannedAll = &v
	return s
}

func (s *DescribeLiveMessageGroupBandResponseBody) SetRequestId(v string) *DescribeLiveMessageGroupBandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveMessageGroupBandResponseBody) SetUnbannedUserList(v []*string) *DescribeLiveMessageGroupBandResponseBody {
	s.UnbannedUserList = v
	return s
}

func (s *DescribeLiveMessageGroupBandResponseBody) Validate() error {
	return dara.Validate(s)
}
