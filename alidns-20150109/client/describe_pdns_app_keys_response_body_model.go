// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKeys(v []*DescribePdnsAppKeysResponseBodyAppKeys) *DescribePdnsAppKeysResponseBody
	GetAppKeys() []*DescribePdnsAppKeysResponseBodyAppKeys
	SetRequestId(v string) *DescribePdnsAppKeysResponseBody
	GetRequestId() *string
}

type DescribePdnsAppKeysResponseBody struct {
	AppKeys   []*DescribePdnsAppKeysResponseBodyAppKeys `json:"AppKeys,omitempty" xml:"AppKeys,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePdnsAppKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeysResponseBody) GetAppKeys() []*DescribePdnsAppKeysResponseBodyAppKeys {
	return s.AppKeys
}

func (s *DescribePdnsAppKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsAppKeysResponseBody) SetAppKeys(v []*DescribePdnsAppKeysResponseBodyAppKeys) *DescribePdnsAppKeysResponseBody {
	s.AppKeys = v
	return s
}

func (s *DescribePdnsAppKeysResponseBody) SetRequestId(v string) *DescribePdnsAppKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBody) Validate() error {
	if s.AppKeys != nil {
		for _, item := range s.AppKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsAppKeysResponseBodyAppKeys struct {
	AppKeyId        *string `json:"AppKeyId,omitempty" xml:"AppKeyId,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePdnsAppKeysResponseBodyAppKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeysResponseBodyAppKeys) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) GetAppKeyId() *string {
	return s.AppKeyId
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) GetCreateDate() *string {
	return s.CreateDate
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) GetRemark() *string {
	return s.Remark
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) GetState() *string {
	return s.State
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) SetAppKeyId(v string) *DescribePdnsAppKeysResponseBodyAppKeys {
	s.AppKeyId = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) SetCreateDate(v string) *DescribePdnsAppKeysResponseBodyAppKeys {
	s.CreateDate = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) SetCreateTimestamp(v int64) *DescribePdnsAppKeysResponseBodyAppKeys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) SetRemark(v string) *DescribePdnsAppKeysResponseBodyAppKeys {
	s.Remark = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) SetState(v string) *DescribePdnsAppKeysResponseBodyAppKeys {
	s.State = &v
	return s
}

func (s *DescribePdnsAppKeysResponseBodyAppKeys) Validate() error {
	return dara.Validate(s)
}
