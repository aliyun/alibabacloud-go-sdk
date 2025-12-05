// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostShareKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKeys(v []*ListHostShareKeysResponseBodyHostShareKeys) *ListHostShareKeysResponseBody
	GetHostShareKeys() []*ListHostShareKeysResponseBodyHostShareKeys
	SetRequestId(v string) *ListHostShareKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListHostShareKeysResponseBody
	GetTotalCount() *int64
}

type ListHostShareKeysResponseBody struct {
	// An array that consists of the shared keys.
	HostShareKeys []*ListHostShareKeysResponseBodyHostShareKeys `json:"HostShareKeys,omitempty" xml:"HostShareKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the shared keys.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostShareKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostShareKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponseBody) GetHostShareKeys() []*ListHostShareKeysResponseBodyHostShareKeys {
	return s.HostShareKeys
}

func (s *ListHostShareKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostShareKeysResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHostShareKeysResponseBody) SetHostShareKeys(v []*ListHostShareKeysResponseBodyHostShareKeys) *ListHostShareKeysResponseBody {
	s.HostShareKeys = v
	return s
}

func (s *ListHostShareKeysResponseBody) SetRequestId(v string) *ListHostShareKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostShareKeysResponseBody) SetTotalCount(v int64) *ListHostShareKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostShareKeysResponseBody) Validate() error {
	if s.HostShareKeys != nil {
		for _, item := range s.HostShareKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHostShareKeysResponseBodyHostShareKeys struct {
	// The number of the associated host accounts.
	//
	// example:
	//
	// 1
	HostAccountCount *int64 `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	// The shared key ID.
	//
	// example:
	//
	// 10247
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The name of the shared key.
	//
	// example:
	//
	// name
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The time when the shared key was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1644806406
	LastModifyKeyAt *int64 `json:"LastModifyKeyAt,omitempty" xml:"LastModifyKeyAt,omitempty"`
	// The fingerprint of the private key.
	//
	// example:
	//
	// ****
	PrivateKeyFingerPrint *string `json:"PrivateKeyFingerPrint,omitempty" xml:"PrivateKeyFingerPrint,omitempty"`
}

func (s ListHostShareKeysResponseBodyHostShareKeys) String() string {
	return dara.Prettify(s)
}

func (s ListHostShareKeysResponseBodyHostShareKeys) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) GetHostAccountCount() *int64 {
	return s.HostAccountCount
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) GetHostShareKeyName() *string {
	return s.HostShareKeyName
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) GetLastModifyKeyAt() *int64 {
	return s.LastModifyKeyAt
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) GetPrivateKeyFingerPrint() *string {
	return s.PrivateKeyFingerPrint
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostAccountCount(v int64) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostAccountCount = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostShareKeyId(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostShareKeyName(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostShareKeyName = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetLastModifyKeyAt(v int64) *ListHostShareKeysResponseBodyHostShareKeys {
	s.LastModifyKeyAt = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetPrivateKeyFingerPrint(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.PrivateKeyFingerPrint = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) Validate() error {
	return dara.Validate(s)
}
