// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportedDataKey(v string) *ExportDataKeyResponseBody
  GetExportedDataKey() *string 
  SetKeyId(v string) *ExportDataKeyResponseBody
  GetKeyId() *string 
  SetKeyVersionId(v string) *ExportDataKeyResponseBody
  GetKeyVersionId() *string 
  SetRequestId(v string) *ExportDataKeyResponseBody
  GetRequestId() *string 
}

type ExportDataKeyResponseBody struct {
  // The data key encrypted by using the public key and then exported.
  // 
  // example:
  // 
  // BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVs*******
  ExportedDataKey *string `json:"ExportedDataKey,omitempty" xml:"ExportedDataKey,omitempty"`
  // The ID of the CMK that is used to decrypt the specified ciphertext of the data key.
  // 
  // This parameter is the globally unique ID of the CMK.
  // 
  // example:
  // 
  // 202b9877-5a25-46e3-a763-e20791b5****
  KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
  // The ID of the CMK version that is used to decrypt the specified ciphertext of the data key.
  // 
  // example:
  // 
  // 2ab1a983-7072-4bbc-a582-584b5bd8****
  KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 4bd560a1-729e-45f1-a3d9-b2a33d61046b
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportDataKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportDataKeyResponseBody) GoString() string {
  return s.String()
}

func (s *ExportDataKeyResponseBody) GetExportedDataKey() *string  {
  return s.ExportedDataKey
}

func (s *ExportDataKeyResponseBody) GetKeyId() *string  {
  return s.KeyId
}

func (s *ExportDataKeyResponseBody) GetKeyVersionId() *string  {
  return s.KeyVersionId
}

func (s *ExportDataKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportDataKeyResponseBody) SetExportedDataKey(v string) *ExportDataKeyResponseBody {
  s.ExportedDataKey = &v
  return s
}

func (s *ExportDataKeyResponseBody) SetKeyId(v string) *ExportDataKeyResponseBody {
  s.KeyId = &v
  return s
}

func (s *ExportDataKeyResponseBody) SetKeyVersionId(v string) *ExportDataKeyResponseBody {
  s.KeyVersionId = &v
  return s
}

func (s *ExportDataKeyResponseBody) SetRequestId(v string) *ExportDataKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportDataKeyResponseBody) Validate() error {
  return dara.Validate(s)
}

