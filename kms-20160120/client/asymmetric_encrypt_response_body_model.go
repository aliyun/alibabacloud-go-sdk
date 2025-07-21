// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricEncryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *AsymmetricEncryptResponseBody
	GetCiphertextBlob() *string
	SetKeyId(v string) *AsymmetricEncryptResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricEncryptResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *AsymmetricEncryptResponseBody
	GetRequestId() *string
}

type AsymmetricEncryptResponseBody struct {
	// The Base64-encoded ciphertext that was generated after encryption.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1Wbjwg==
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AsymmetricEncryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *AsymmetricEncryptResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricEncryptResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricEncryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsymmetricEncryptResponseBody) SetCiphertextBlob(v string) *AsymmetricEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetKeyId(v string) *AsymmetricEncryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetKeyVersionId(v string) *AsymmetricEncryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) SetRequestId(v string) *AsymmetricEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricEncryptResponseBody) Validate() error {
	return dara.Validate(s)
}
