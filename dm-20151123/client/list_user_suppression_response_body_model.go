// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSuppressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListUserSuppressionResponseBodyData) *ListUserSuppressionResponseBody
	GetData() *ListUserSuppressionResponseBodyData
	SetPageNumber(v int32) *ListUserSuppressionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserSuppressionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserSuppressionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserSuppressionResponseBody
	GetTotalCount() *int32
}

type ListUserSuppressionResponseBody struct {
	// Returned results.
	Data *ListUserSuppressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 1A846D66-5EC7-551B-9687-5BF1963DCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserSuppressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBody) GetData() *ListUserSuppressionResponseBodyData {
	return s.Data
}

func (s *ListUserSuppressionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserSuppressionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSuppressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserSuppressionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserSuppressionResponseBody) SetData(v *ListUserSuppressionResponseBodyData) *ListUserSuppressionResponseBody {
	s.Data = v
	return s
}

func (s *ListUserSuppressionResponseBody) SetPageNumber(v int32) *ListUserSuppressionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetPageSize(v int32) *ListUserSuppressionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetRequestId(v string) *ListUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetTotalCount(v int32) *ListUserSuppressionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserSuppressionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserSuppressionResponseBodyData struct {
	UserSuppressions []*ListUserSuppressionResponseBodyDataUserSuppressions `json:"UserSuppressions,omitempty" xml:"UserSuppressions,omitempty" type:"Repeated"`
}

func (s ListUserSuppressionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserSuppressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBodyData) GetUserSuppressions() []*ListUserSuppressionResponseBodyDataUserSuppressions {
	return s.UserSuppressions
}

func (s *ListUserSuppressionResponseBodyData) SetUserSuppressions(v []*ListUserSuppressionResponseBodyDataUserSuppressions) *ListUserSuppressionResponseBodyData {
	s.UserSuppressions = v
	return s
}

func (s *ListUserSuppressionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUserSuppressionResponseBodyDataUserSuppressions struct {
	// Email address or domain name
	//
	// example:
	//
	// test@example.net
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Creation time, timestamp, accurate to the second.
	//
	// example:
	//
	// 1715667435
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Last bounce hit time, timestamp, accurate to the second.
	//
	// example:
	//
	// 1715667451
	LastBounceTime *int32 `json:"LastBounceTime,omitempty" xml:"LastBounceTime,omitempty"`
	// Invalid address ID
	//
	// example:
	//
	// 59511
	SuppressionId *int32 `json:"SuppressionId,omitempty" xml:"SuppressionId,omitempty"`
	// Source of entry, invalid address type
	//
	// - system
	//
	// - user
	//
	// example:
	//
	// user
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUserSuppressionResponseBodyDataUserSuppressions) String() string {
	return dara.Prettify(s)
}

func (s ListUserSuppressionResponseBodyDataUserSuppressions) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) GetAddress() *string {
	return s.Address
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) GetLastBounceTime() *int32 {
	return s.LastBounceTime
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) GetSuppressionId() *int32 {
	return s.SuppressionId
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) GetType() *string {
	return s.Type
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetAddress(v string) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.Address = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetCreateTime(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.CreateTime = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetLastBounceTime(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.LastBounceTime = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetSuppressionId(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.SuppressionId = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetType(v string) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.Type = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) Validate() error {
	return dara.Validate(s)
}
