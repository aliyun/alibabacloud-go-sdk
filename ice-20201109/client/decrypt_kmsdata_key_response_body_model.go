// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptKMSDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataKey(v *DecryptKMSDataKeyResponseBodyDataKey) *DecryptKMSDataKeyResponseBody
	GetDataKey() *DecryptKMSDataKeyResponseBodyDataKey
	SetRequestId(v string) *DecryptKMSDataKeyResponseBody
	GetRequestId() *string
}

type DecryptKMSDataKeyResponseBody struct {
	// The information about the decryption result.
	DataKey *DecryptKMSDataKeyResponseBodyDataKey `json:"DataKey,omitempty" xml:"DataKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptKMSDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyResponseBody) GetDataKey() *DecryptKMSDataKeyResponseBodyDataKey {
	return s.DataKey
}

func (s *DecryptKMSDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecryptKMSDataKeyResponseBody) SetDataKey(v *DecryptKMSDataKeyResponseBodyDataKey) *DecryptKMSDataKeyResponseBody {
	s.DataKey = v
	return s
}

func (s *DecryptKMSDataKeyResponseBody) SetRequestId(v string) *DecryptKMSDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DecryptKMSDataKeyResponseBodyDataKey struct {
	// The ID of the customer master key (CMK) that was used to decrypt the ciphertext.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext that is generated after decryption.
	//
	// example:
	//
	// tRYXuCwgja12xxO1N/gZERDDCLw9doZEQiPDk/Bv****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s DecryptKMSDataKeyResponseBodyDataKey) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyResponseBodyDataKey) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyResponseBodyDataKey) GetKeyId() *string {
	return s.KeyId
}

func (s *DecryptKMSDataKeyResponseBodyDataKey) GetPlaintext() *string {
	return s.Plaintext
}

func (s *DecryptKMSDataKeyResponseBodyDataKey) SetKeyId(v string) *DecryptKMSDataKeyResponseBodyDataKey {
	s.KeyId = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBodyDataKey) SetPlaintext(v string) *DecryptKMSDataKeyResponseBodyDataKey {
	s.Plaintext = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBodyDataKey) Validate() error {
	return dara.Validate(s)
}
