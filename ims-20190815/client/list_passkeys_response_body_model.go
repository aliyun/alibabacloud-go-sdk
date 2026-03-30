// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPasskeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasskeys(v []*ListPasskeysResponseBodyPasskeys) *ListPasskeysResponseBody
	GetPasskeys() []*ListPasskeysResponseBodyPasskeys
	SetRequestId(v string) *ListPasskeysResponseBody
	GetRequestId() *string
}

type ListPasskeysResponseBody struct {
	// The information about the passkeys.
	Passkeys []*ListPasskeysResponseBodyPasskeys `json:"Passkeys,omitempty" xml:"Passkeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPasskeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPasskeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponseBody) GetPasskeys() []*ListPasskeysResponseBodyPasskeys {
	return s.Passkeys
}

func (s *ListPasskeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPasskeysResponseBody) SetPasskeys(v []*ListPasskeysResponseBodyPasskeys) *ListPasskeysResponseBody {
	s.Passkeys = v
	return s
}

func (s *ListPasskeysResponseBody) SetRequestId(v string) *ListPasskeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPasskeysResponseBody) Validate() error {
	if s.Passkeys != nil {
		for _, item := range s.Passkeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPasskeysResponseBodyPasskeys struct {
	// The time when the passkey was created. The value is a timestamp.
	//
	// example:
	//
	// 1737450279
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the passkey was last used. The value is a timestamp.
	//
	// example:
	//
	// 1737450390
	LastUseDate *string `json:"LastUseDate,omitempty" xml:"LastUseDate,omitempty"`
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The name of the passkey.
	//
	// example:
	//
	// device1
	PasskeyName *string `json:"PasskeyName,omitempty" xml:"PasskeyName,omitempty"`
}

func (s ListPasskeysResponseBodyPasskeys) String() string {
	return dara.Prettify(s)
}

func (s ListPasskeysResponseBodyPasskeys) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponseBodyPasskeys) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListPasskeysResponseBodyPasskeys) GetLastUseDate() *string {
	return s.LastUseDate
}

func (s *ListPasskeysResponseBodyPasskeys) GetPasskeyId() *string {
	return s.PasskeyId
}

func (s *ListPasskeysResponseBodyPasskeys) GetPasskeyName() *string {
	return s.PasskeyName
}

func (s *ListPasskeysResponseBodyPasskeys) SetCreateDate(v string) *ListPasskeysResponseBodyPasskeys {
	s.CreateDate = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetLastUseDate(v string) *ListPasskeysResponseBodyPasskeys {
	s.LastUseDate = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetPasskeyId(v string) *ListPasskeysResponseBodyPasskeys {
	s.PasskeyId = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetPasskeyName(v string) *ListPasskeysResponseBodyPasskeys {
	s.PasskeyName = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) Validate() error {
	return dara.Validate(s)
}
