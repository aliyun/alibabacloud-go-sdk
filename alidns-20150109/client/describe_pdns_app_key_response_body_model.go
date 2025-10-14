// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v *DescribePdnsAppKeyResponseBodyAppKey) *DescribePdnsAppKeyResponseBody
	GetAppKey() *DescribePdnsAppKeyResponseBodyAppKey
	SetRequestId(v string) *DescribePdnsAppKeyResponseBody
	GetRequestId() *string
}

type DescribePdnsAppKeyResponseBody struct {
	AppKey    *DescribePdnsAppKeyResponseBodyAppKey `json:"AppKey,omitempty" xml:"AppKey,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePdnsAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeyResponseBody) GetAppKey() *DescribePdnsAppKeyResponseBodyAppKey {
	return s.AppKey
}

func (s *DescribePdnsAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsAppKeyResponseBody) SetAppKey(v *DescribePdnsAppKeyResponseBodyAppKey) *DescribePdnsAppKeyResponseBody {
	s.AppKey = v
	return s
}

func (s *DescribePdnsAppKeyResponseBody) SetRequestId(v string) *DescribePdnsAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBody) Validate() error {
	if s.AppKey != nil {
		if err := s.AppKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePdnsAppKeyResponseBodyAppKey struct {
	AppKeyId        *string `json:"AppKeyId,omitempty" xml:"AppKeyId,omitempty"`
	AppKeySecret    *string `json:"AppKeySecret,omitempty" xml:"AppKeySecret,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePdnsAppKeyResponseBodyAppKey) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAppKeyResponseBodyAppKey) GoString() string {
	return s.String()
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetAppKeyId() *string {
	return s.AppKeyId
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetAppKeySecret() *string {
	return s.AppKeySecret
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetCreateDate() *string {
	return s.CreateDate
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetRemark() *string {
	return s.Remark
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) GetState() *string {
	return s.State
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetAppKeyId(v string) *DescribePdnsAppKeyResponseBodyAppKey {
	s.AppKeyId = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetAppKeySecret(v string) *DescribePdnsAppKeyResponseBodyAppKey {
	s.AppKeySecret = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetCreateDate(v string) *DescribePdnsAppKeyResponseBodyAppKey {
	s.CreateDate = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetCreateTimestamp(v int64) *DescribePdnsAppKeyResponseBodyAppKey {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetRemark(v string) *DescribePdnsAppKeyResponseBodyAppKey {
	s.Remark = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) SetState(v string) *DescribePdnsAppKeyResponseBodyAppKey {
	s.State = &v
	return s
}

func (s *DescribePdnsAppKeyResponseBodyAppKey) Validate() error {
	return dara.Validate(s)
}
