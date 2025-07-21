// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvalidAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextStart(v string) *QueryInvalidAddressResponseBody
	GetNextStart() *string
	SetRequestId(v string) *QueryInvalidAddressResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryInvalidAddressResponseBody
	GetTotalCount() *int32
	SetData(v *QueryInvalidAddressResponseBodyData) *QueryInvalidAddressResponseBody
	GetData() *QueryInvalidAddressResponseBodyData
}

type QueryInvalidAddressResponseBody struct {
	// Next request starting position.
	//
	// example:
	//
	// 2
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Records.
	Data *QueryInvalidAddressResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryInvalidAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBody) GetNextStart() *string {
	return s.NextStart
}

func (s *QueryInvalidAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInvalidAddressResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryInvalidAddressResponseBody) GetData() *QueryInvalidAddressResponseBodyData {
	return s.Data
}

func (s *QueryInvalidAddressResponseBody) SetNextStart(v string) *QueryInvalidAddressResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetRequestId(v string) *QueryInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetTotalCount(v int32) *QueryInvalidAddressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetData(v *QueryInvalidAddressResponseBodyData) *QueryInvalidAddressResponseBody {
	s.Data = v
	return s
}

func (s *QueryInvalidAddressResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryInvalidAddressResponseBodyData struct {
	MailDetail []*QueryInvalidAddressResponseBodyDataMailDetail `json:"mailDetail,omitempty" xml:"mailDetail,omitempty" type:"Repeated"`
}

func (s QueryInvalidAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInvalidAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBodyData) GetMailDetail() []*QueryInvalidAddressResponseBodyDataMailDetail {
	return s.MailDetail
}

func (s *QueryInvalidAddressResponseBodyData) SetMailDetail(v []*QueryInvalidAddressResponseBodyDataMailDetail) *QueryInvalidAddressResponseBodyData {
	s.MailDetail = v
	return s
}

func (s *QueryInvalidAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryInvalidAddressResponseBodyDataMailDetail struct {
	// Update time.
	//
	// example:
	//
	// 2021-04-28T17:11Z
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// Recipient address.
	//
	// example:
	//
	// toaddress@example.com
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// Update time (in timestamp format).
	//
	// example:
	//
	// 1619601108
	UtcLastUpdateTime *int64 `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
}

func (s QueryInvalidAddressResponseBodyDataMailDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryInvalidAddressResponseBodyDataMailDetail) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) GetToAddress() *string {
	return s.ToAddress
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) GetUtcLastUpdateTime() *int64 {
	return s.UtcLastUpdateTime
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetLastUpdateTime(v string) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetToAddress(v string) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.ToAddress = &v
	return s
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetUtcLastUpdateTime(v int64) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.UtcLastUpdateTime = &v
	return s
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) Validate() error {
	return dara.Validate(s)
}
