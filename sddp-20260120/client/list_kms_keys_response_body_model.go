// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListKmsKeysResponseBodyItems) *ListKmsKeysResponseBody
	GetItems() []*ListKmsKeysResponseBodyItems
	SetRequestId(v string) *ListKmsKeysResponseBody
	GetRequestId() *string
}

type ListKmsKeysResponseBody struct {
	Items     []*ListKmsKeysResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListKmsKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKmsKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListKmsKeysResponseBody) GetItems() []*ListKmsKeysResponseBodyItems {
	return s.Items
}

func (s *ListKmsKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKmsKeysResponseBody) SetItems(v []*ListKmsKeysResponseBodyItems) *ListKmsKeysResponseBody {
	s.Items = v
	return s
}

func (s *ListKmsKeysResponseBody) SetRequestId(v string) *ListKmsKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKmsKeysResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKmsKeysResponseBodyItems struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListKmsKeysResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListKmsKeysResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListKmsKeysResponseBodyItems) GetAliasName() *string {
	return s.AliasName
}

func (s *ListKmsKeysResponseBodyItems) GetKeyId() *string {
	return s.KeyId
}

func (s *ListKmsKeysResponseBodyItems) SetAliasName(v string) *ListKmsKeysResponseBodyItems {
	s.AliasName = &v
	return s
}

func (s *ListKmsKeysResponseBodyItems) SetKeyId(v string) *ListKmsKeysResponseBodyItems {
	s.KeyId = &v
	return s
}

func (s *ListKmsKeysResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
