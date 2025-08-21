// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeys(v *ListAccessKeysResponseBodyAccessKeys) *ListAccessKeysResponseBody
	GetAccessKeys() *ListAccessKeysResponseBodyAccessKeys
	SetRequestId(v string) *ListAccessKeysResponseBody
	GetRequestId() *string
}

type ListAccessKeysResponseBody struct {
	// The AccessKey pairs.
	AccessKeys *ListAccessKeysResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBody) GetAccessKeys() *ListAccessKeysResponseBodyAccessKeys {
	return s.AccessKeys
}

func (s *ListAccessKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessKeysResponseBody) SetAccessKeys(v *ListAccessKeysResponseBodyAccessKeys) *ListAccessKeysResponseBody {
	s.AccessKeys = v
	return s
}

func (s *ListAccessKeysResponseBody) SetRequestId(v string) *ListAccessKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAccessKeysResponseBodyAccessKeys struct {
	AccessKey []*ListAccessKeysResponseBodyAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Repeated"`
}

func (s ListAccessKeysResponseBodyAccessKeys) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeys) GetAccessKey() []*ListAccessKeysResponseBodyAccessKeysAccessKey {
	return s.AccessKey
}

func (s *ListAccessKeysResponseBodyAccessKeys) SetAccessKey(v []*ListAccessKeysResponseBodyAccessKeysAccessKey) *ListAccessKeysResponseBodyAccessKeys {
	s.AccessKey = v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeys) Validate() error {
	return dara.Validate(s)
}

type ListAccessKeysResponseBodyAccessKeysAccessKey struct {
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
	// 2020-10-13T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the AccessKey pair was updated.
	//
	// example:
	//
	// 2020-10-13T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) GetStatus() *string {
	return s.Status
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetStatus(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.Status = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetUpdateDate(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.UpdateDate = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) Validate() error {
	return dara.Validate(s)
}
