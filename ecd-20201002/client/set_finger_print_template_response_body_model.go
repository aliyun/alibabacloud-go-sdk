// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedPassword(v string) *SetFingerPrintTemplateResponseBody
	GetEncryptedPassword() *string
	SetIndex(v int32) *SetFingerPrintTemplateResponseBody
	GetIndex() *int32
	SetRequestId(v string) *SetFingerPrintTemplateResponseBody
	GetRequestId() *string
}

type SetFingerPrintTemplateResponseBody struct {
	// The encrypted password.
	//
	// example:
	//
	// 0711abb9-4cf8-41b2-9d0e-b51209468631;da4VFPNxwY3CZegFjOrCNw==;iHp2l9/qGcfD4tWx7jZIZQ==
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	// The index.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CDE666EA-4FCD-5024-895C-8698E3D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponseBody) GetEncryptedPassword() *string {
	return s.EncryptedPassword
}

func (s *SetFingerPrintTemplateResponseBody) GetIndex() *int32 {
	return s.Index
}

func (s *SetFingerPrintTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetFingerPrintTemplateResponseBody) SetEncryptedPassword(v string) *SetFingerPrintTemplateResponseBody {
	s.EncryptedPassword = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetIndex(v int32) *SetFingerPrintTemplateResponseBody {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetRequestId(v string) *SetFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
