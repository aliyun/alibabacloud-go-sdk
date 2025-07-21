// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAndExportDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *GenerateAndExportDataKeyResponseBody
	GetCiphertextBlob() *string
	SetExportedDataKey(v string) *GenerateAndExportDataKeyResponseBody
	GetExportedDataKey() *string
	SetKeyId(v string) *GenerateAndExportDataKeyResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *GenerateAndExportDataKeyResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *GenerateAndExportDataKeyResponseBody
	GetRequestId() *string
}

type GenerateAndExportDataKeyResponseBody struct {
	// The ciphertext of the data key encrypted by using the primary CMK version.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The data key encrypted by using the public key and then exported.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVs*******
	ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 599fa825-17de-417e-9554-bb032cc6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to encrypt the plaintext. It is the primary version of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAndExportDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAndExportDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *GenerateAndExportDataKeyResponseBody) GetExportedDataKey() *string {
	return s.ExportedDataKey
}

func (s *GenerateAndExportDataKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateAndExportDataKeyResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *GenerateAndExportDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAndExportDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateAndExportDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetExportedDataKey(v string) *GenerateAndExportDataKeyResponseBody {
	s.ExportedDataKey = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetKeyId(v string) *GenerateAndExportDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetKeyVersionId(v string) *GenerateAndExportDataKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) SetRequestId(v string) *GenerateAndExportDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAndExportDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
