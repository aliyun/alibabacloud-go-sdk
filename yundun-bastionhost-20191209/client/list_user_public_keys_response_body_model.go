// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPublicKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicKeys(v []*ListUserPublicKeysResponseBodyPublicKeys) *ListUserPublicKeysResponseBody
	GetPublicKeys() []*ListUserPublicKeysResponseBodyPublicKeys
	SetRequestId(v string) *ListUserPublicKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUserPublicKeysResponseBody
	GetTotalCount() *int64
}

type ListUserPublicKeysResponseBody struct {
	// An array that consists of the public keys of the user.
	PublicKeys []*ListUserPublicKeysResponseBodyPublicKeys `json:"PublicKeys,omitempty" xml:"PublicKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of public keys.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserPublicKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPublicKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponseBody) GetPublicKeys() []*ListUserPublicKeysResponseBodyPublicKeys {
	return s.PublicKeys
}

func (s *ListUserPublicKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPublicKeysResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserPublicKeysResponseBody) SetPublicKeys(v []*ListUserPublicKeysResponseBodyPublicKeys) *ListUserPublicKeysResponseBody {
	s.PublicKeys = v
	return s
}

func (s *ListUserPublicKeysResponseBody) SetRequestId(v string) *ListUserPublicKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPublicKeysResponseBody) SetTotalCount(v int64) *ListUserPublicKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPublicKeysResponseBody) Validate() error {
	if s.PublicKeys != nil {
		for _, item := range s.PublicKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPublicKeysResponseBodyPublicKeys struct {
	// The description of the public key.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The fingerprint of the public key.
	//
	// example:
	//
	// d8:7d:b6:27:70:2d:07:fb:c6:b6:66:0a:86:7b:0f:9a
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	// The ID of the public key.
	//
	// example:
	//
	// 1
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The name of the public key.
	//
	// example:
	//
	// Keyname
	PublicKeyName *string `json:"PublicKeyName,omitempty" xml:"PublicKeyName,omitempty"`
	// The ID of the user to which the public key belongs.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserPublicKeysResponseBodyPublicKeys) String() string {
	return dara.Prettify(s)
}

func (s ListUserPublicKeysResponseBodyPublicKeys) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) GetComment() *string {
	return s.Comment
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) GetPublicKeyId() *string {
	return s.PublicKeyId
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) GetPublicKeyName() *string {
	return s.PublicKeyName
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) GetUserId() *string {
	return s.UserId
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetComment(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.Comment = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetFingerPrint(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.FingerPrint = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetPublicKeyId(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.PublicKeyId = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetPublicKeyName(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.PublicKeyName = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetUserId(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.UserId = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) Validate() error {
	return dara.Validate(s)
}
