// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKey(v *GetHostShareKeyResponseBodyHostShareKey) *GetHostShareKeyResponseBody
	GetHostShareKey() *GetHostShareKeyResponseBodyHostShareKey
	SetRequestId(v string) *GetHostShareKeyResponseBody
	GetRequestId() *string
}

type GetHostShareKeyResponseBody struct {
	// The returned information about the shared key.
	HostShareKey *GetHostShareKeyResponseBodyHostShareKey `json:"HostShareKey,omitempty" xml:"HostShareKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponseBody) GetHostShareKey() *GetHostShareKeyResponseBodyHostShareKey {
	return s.HostShareKey
}

func (s *GetHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostShareKeyResponseBody) SetHostShareKey(v *GetHostShareKeyResponseBodyHostShareKey) *GetHostShareKeyResponseBody {
	s.HostShareKey = v
	return s
}

func (s *GetHostShareKeyResponseBody) SetRequestId(v string) *GetHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostShareKeyResponseBody) Validate() error {
	if s.HostShareKey != nil {
		if err := s.HostShareKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHostShareKeyResponseBodyHostShareKey struct {
	// The ID of the shared key.
	//
	// example:
	//
	// 10427
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The name of the shared key.
	//
	// example:
	//
	// name
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The time when the information about the shared key was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1644287246
	LastModifyKeyAt *int64 `json:"LastModifyKeyAt,omitempty" xml:"LastModifyKeyAt,omitempty"`
	// The fingerprint of the private key.
	//
	// example:
	//
	// ***
	PrivateKeyFingerPrint *string `json:"PrivateKeyFingerPrint,omitempty" xml:"PrivateKeyFingerPrint,omitempty"`
}

func (s GetHostShareKeyResponseBodyHostShareKey) String() string {
	return dara.Prettify(s)
}

func (s GetHostShareKeyResponseBodyHostShareKey) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponseBodyHostShareKey) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *GetHostShareKeyResponseBodyHostShareKey) GetHostShareKeyName() *string {
	return s.HostShareKeyName
}

func (s *GetHostShareKeyResponseBodyHostShareKey) GetLastModifyKeyAt() *int64 {
	return s.LastModifyKeyAt
}

func (s *GetHostShareKeyResponseBodyHostShareKey) GetPrivateKeyFingerPrint() *string {
	return s.PrivateKeyFingerPrint
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetHostShareKeyId(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetHostShareKeyName(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.HostShareKeyName = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetLastModifyKeyAt(v int64) *GetHostShareKeyResponseBodyHostShareKey {
	s.LastModifyKeyAt = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetPrivateKeyFingerPrint(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.PrivateKeyFingerPrint = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) Validate() error {
	return dara.Validate(s)
}
