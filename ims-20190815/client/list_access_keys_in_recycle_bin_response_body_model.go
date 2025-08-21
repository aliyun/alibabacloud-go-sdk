// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeys(v *ListAccessKeysInRecycleBinResponseBodyAccessKeys) *ListAccessKeysInRecycleBinResponseBody
	GetAccessKeys() *ListAccessKeysInRecycleBinResponseBodyAccessKeys
	SetRequestId(v string) *ListAccessKeysInRecycleBinResponseBody
	GetRequestId() *string
}

type ListAccessKeysInRecycleBinResponseBody struct {
	// The information about the AccessKey pairs.
	AccessKeys *ListAccessKeysInRecycleBinResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBody) GetAccessKeys() *ListAccessKeysInRecycleBinResponseBodyAccessKeys {
	return s.AccessKeys
}

func (s *ListAccessKeysInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessKeysInRecycleBinResponseBody) SetAccessKeys(v *ListAccessKeysInRecycleBinResponseBodyAccessKeys) *ListAccessKeysInRecycleBinResponseBody {
	s.AccessKeys = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBody) SetRequestId(v string) *ListAccessKeysInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAccessKeysInRecycleBinResponseBodyAccessKeys struct {
	AccessKey []*ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Repeated"`
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeys) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeys) GetAccessKey() []*ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	return s.AccessKey
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeys) SetAccessKey(v []*ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) *ListAccessKeysInRecycleBinResponseBodyAccessKeys {
	s.AccessKey = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeys) Validate() error {
	return dara.Validate(s)
}

type ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-11T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the AccessKey pair will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-12T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The time when the AccessKey pair was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GetRecycleDate() *string {
	return s.RecycleDate
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetDeleteDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.DeleteDate = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetRecycleDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.RecycleDate = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) Validate() error {
	return dara.Validate(s)
}
