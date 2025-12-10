// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeyDeliveriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAICPublicKeyDeliveriesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAICPublicKeyDeliveriesResponseBody
	GetPageSize() *int64
	SetPublicKeyDeliverInfo(v []*ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) *ListAICPublicKeyDeliveriesResponseBody
	GetPublicKeyDeliverInfo() []*ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo
	SetRequestId(v string) *ListAICPublicKeyDeliveriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAICPublicKeyDeliveriesResponseBody
	GetTotalCount() *int64
}

type ListAICPublicKeyDeliveriesResponseBody struct {
	// Current page number when paging
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Public Key List
	PublicKeyDeliverInfo []*ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo `json:"PublicKeyDeliverInfo,omitempty" xml:"PublicKeyDeliverInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// xxxsxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAICPublicKeyDeliveriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeyDeliveriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeyDeliveriesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAICPublicKeyDeliveriesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAICPublicKeyDeliveriesResponseBody) GetPublicKeyDeliverInfo() []*ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	return s.PublicKeyDeliverInfo
}

func (s *ListAICPublicKeyDeliveriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAICPublicKeyDeliveriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAICPublicKeyDeliveriesResponseBody) SetPageNumber(v int64) *ListAICPublicKeyDeliveriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBody) SetPageSize(v int64) *ListAICPublicKeyDeliveriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBody) SetPublicKeyDeliverInfo(v []*ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) *ListAICPublicKeyDeliveriesResponseBody {
	s.PublicKeyDeliverInfo = v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBody) SetRequestId(v string) *ListAICPublicKeyDeliveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBody) SetTotalCount(v int64) *ListAICPublicKeyDeliveriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBody) Validate() error {
	if s.PublicKeyDeliverInfo != nil {
		for _, item := range s.PublicKeyDeliverInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo struct {
	// The creation time.
	//
	// example:
	//
	// 2025-10-09T15:13:21+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// aic-xxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Public Key Grouping
	//
	// example:
	//
	// test-group
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// Public Key Name
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// Public key type
	//
	// example:
	//
	// adb
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) GetKeyType() *string {
	return s.KeyType
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) SetCreationTime(v string) *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	s.CreationTime = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) SetInstanceId(v string) *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	s.InstanceId = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) SetKeyGroup(v string) *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	s.KeyGroup = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) SetKeyName(v string) *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	s.KeyName = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) SetKeyType(v string) *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo {
	s.KeyType = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponseBodyPublicKeyDeliverInfo) Validate() error {
	return dara.Validate(s)
}
