// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *DecryptRequest
	GetCiphertextBlob() *string
	SetDryRun(v string) *DecryptRequest
	GetDryRun() *string
	SetEncryptionContext(v map[string]interface{}) *DecryptRequest
	GetEncryptionContext() map[string]interface{}
}

type DecryptRequest struct {
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
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s DecryptRequest) String() string {
	return dara.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *DecryptRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *DecryptRequest) GetEncryptionContext() map[string]interface{} {
	return s.EncryptionContext
}

func (s *DecryptRequest) SetCiphertextBlob(v string) *DecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptRequest) SetDryRun(v string) *DecryptRequest {
	s.DryRun = &v
	return s
}

func (s *DecryptRequest) SetEncryptionContext(v map[string]interface{}) *DecryptRequest {
	s.EncryptionContext = v
	return s
}

func (s *DecryptRequest) Validate() error {
	return dara.Validate(s)
}
