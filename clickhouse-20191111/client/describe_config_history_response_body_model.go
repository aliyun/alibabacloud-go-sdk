// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigHistoryItems(v []*DescribeConfigHistoryResponseBodyConfigHistoryItems) *DescribeConfigHistoryResponseBody
	GetConfigHistoryItems() []*DescribeConfigHistoryResponseBodyConfigHistoryItems
	SetRequestId(v string) *DescribeConfigHistoryResponseBody
	GetRequestId() *string
}

type DescribeConfigHistoryResponseBody struct {
	// The change records of the configuration parameters.
	ConfigHistoryItems []*DescribeConfigHistoryResponseBodyConfigHistoryItems `json:"ConfigHistoryItems,omitempty" xml:"ConfigHistoryItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBody) GetConfigHistoryItems() []*DescribeConfigHistoryResponseBodyConfigHistoryItems {
	return s.ConfigHistoryItems
}

func (s *DescribeConfigHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigHistoryResponseBody) SetConfigHistoryItems(v []*DescribeConfigHistoryResponseBodyConfigHistoryItems) *DescribeConfigHistoryResponseBody {
	s.ConfigHistoryItems = v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetRequestId(v string) *DescribeConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBody) Validate() error {
	if s.ConfigHistoryItems != nil {
		for _, item := range s.ConfigHistoryItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfigHistoryResponseBodyConfigHistoryItems struct {
	// The ID of the change record.
	//
	// example:
	//
	// 1
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	// The user ID (UID) of the Alibaba Cloud account.
	//
	// example:
	//
	// 253460731706911258
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reason for the setting modification of the configuration parameters.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether the setting modification of the configuration parameters took effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the values of the configuration parameters were changed. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-22T10:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeConfigHistoryResponseBodyConfigHistoryItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponseBodyConfigHistoryItems) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) GetChangeId() *string {
	return s.ChangeId
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) GetReason() *string {
	return s.Reason
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) GetTime() *string {
	return s.Time
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetChangeId(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.ChangeId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetOwnerId(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.OwnerId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetReason(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Reason = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetSuccess(v bool) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Success = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetTime(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Time = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) Validate() error {
	return dara.Validate(s)
}
