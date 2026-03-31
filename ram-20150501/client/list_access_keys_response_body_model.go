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
	if s.AccessKeys != nil {
		if err := s.AccessKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.AccessKey != nil {
		for _, item := range s.AccessKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessKeysResponseBodyAccessKeysAccessKey struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) Validate() error {
	return dara.Validate(s)
}
