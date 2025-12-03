// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeColdStorageResponseBody
	GetClusterId() *string
	SetColdStorageSize(v string) *DescribeColdStorageResponseBody
	GetColdStorageSize() *string
	SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody
	GetColdStorageUsePercent() *string
	SetOpenStatus(v string) *DescribeColdStorageResponseBody
	GetOpenStatus() *string
	SetPayType(v string) *DescribeColdStorageResponseBody
	GetPayType() *string
	SetRequestId(v string) *DescribeColdStorageResponseBody
	GetRequestId() *string
}

type DescribeColdStorageResponseBody struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ColdStorageSize       *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	ColdStorageUsePercent *string `json:"ColdStorageUsePercent,omitempty" xml:"ColdStorageUsePercent,omitempty"`
	OpenStatus            *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeColdStorageResponseBody) GetColdStorageSize() *string {
	return s.ColdStorageSize
}

func (s *DescribeColdStorageResponseBody) GetColdStorageUsePercent() *string {
	return s.ColdStorageUsePercent
}

func (s *DescribeColdStorageResponseBody) GetOpenStatus() *string {
	return s.OpenStatus
}

func (s *DescribeColdStorageResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeColdStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColdStorageResponseBody) SetClusterId(v string) *DescribeColdStorageResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageSize(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUsePercent = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetOpenStatus(v string) *DescribeColdStorageResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetPayType(v string) *DescribeColdStorageResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetRequestId(v string) *DescribeColdStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
