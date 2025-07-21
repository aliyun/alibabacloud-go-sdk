// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *DecryptShrinkRequest
	GetCiphertextBlob() *string
	SetDryRun(v string) *DecryptShrinkRequest
	GetDryRun() *string
	SetEncryptionContextShrink(v string) *DecryptShrinkRequest
	GetEncryptionContextShrink() *string
}

type DecryptShrinkRequest struct {
	// The ciphertext that you want to decrypt.
	//
	// You can generate the ciphertext by calling the following operations:
	//
	// 	- [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html)
	//
	// 	- [Encrypt](https://help.aliyun.com/document_detail/28949.html)
	//
	// 	- [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	DryRun         *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// >  If you specify the EncryptionContext parameter when you call the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [Encrypt](https://help.aliyun.com/document_detail/28949.html), or [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html) operation, you must specify the same context when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s DecryptShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DecryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *DecryptShrinkRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *DecryptShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *DecryptShrinkRequest) GetEncryptionContextShrink() *string {
	return s.EncryptionContextShrink
}

func (s *DecryptShrinkRequest) SetCiphertextBlob(v string) *DecryptShrinkRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptShrinkRequest) SetDryRun(v string) *DecryptShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *DecryptShrinkRequest) SetEncryptionContextShrink(v string) *DecryptShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *DecryptShrinkRequest) Validate() error {
	return dara.Validate(s)
}
