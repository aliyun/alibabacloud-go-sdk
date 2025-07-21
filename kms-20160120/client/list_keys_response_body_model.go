// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v *ListKeysResponseBodyKeys) *ListKeysResponseBody
	GetKeys() *ListKeysResponseBodyKeys
	SetPageNumber(v int32) *ListKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKeysResponseBody
	GetTotalCount() *int32
}

type ListKeysResponseBody struct {
	// An array that consists of the CMKs of the current Alibaba Cloud account in the current region.
	Keys *ListKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8252db58-2036-408c-a3d5-56e656dc2551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of CMKs.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBody) GetKeys() *ListKeysResponseBodyKeys {
	return s.Keys
}

func (s *ListKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKeysResponseBody) SetKeys(v *ListKeysResponseBodyKeys) *ListKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListKeysResponseBody) SetPageNumber(v int32) *ListKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeysResponseBody) SetPageSize(v int32) *ListKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeysResponseBody) SetRequestId(v string) *ListKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeysResponseBody) SetTotalCount(v int32) *ListKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListKeysResponseBodyKeys struct {
	Key []*ListKeysResponseBodyKeysKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListKeysResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBodyKeys) GetKey() []*ListKeysResponseBodyKeysKey {
	return s.Key
}

func (s *ListKeysResponseBodyKeys) SetKey(v []*ListKeysResponseBodyKeysKey) *ListKeysResponseBodyKeys {
	s.Key = v
	return s
}

func (s *ListKeysResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

type ListKeysResponseBodyKeysKey struct {
	// The Alibaba Cloud Resource Name (ARN) of the CMK.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:123456:key/80e9409f-78fa-42ab-84bd-83f40c81****
	KeyArn *string `json:"KeyArn,omitempty" xml:"KeyArn,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListKeysResponseBodyKeysKey) String() string {
	return dara.Prettify(s)
}

func (s ListKeysResponseBodyKeysKey) GoString() string {
	return s.String()
}

func (s *ListKeysResponseBodyKeysKey) GetKeyArn() *string {
	return s.KeyArn
}

func (s *ListKeysResponseBodyKeysKey) GetKeyId() *string {
	return s.KeyId
}

func (s *ListKeysResponseBodyKeysKey) SetKeyArn(v string) *ListKeysResponseBodyKeysKey {
	s.KeyArn = &v
	return s
}

func (s *ListKeysResponseBodyKeysKey) SetKeyId(v string) *ListKeysResponseBodyKeysKey {
	s.KeyId = &v
	return s
}

func (s *ListKeysResponseBodyKeysKey) Validate() error {
	return dara.Validate(s)
}
