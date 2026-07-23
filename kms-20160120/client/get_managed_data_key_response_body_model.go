// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataKeyName(v string) *GetManagedDataKeyResponseBody
	GetDataKeyName() *string
	SetDataKeyVersionId(v string) *GetManagedDataKeyResponseBody
	GetDataKeyVersionId() *string
	SetDataKeyVersionName(v string) *GetManagedDataKeyResponseBody
	GetDataKeyVersionName() *string
	SetPlaintext(v string) *GetManagedDataKeyResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *GetManagedDataKeyResponseBody
	GetRequestId() *string
}

type GetManagedDataKeyResponseBody struct {
	// The name of the managed data key (DK).
	//
	// example:
	//
	// example-data-key
	DataKeyName *string `json:"DataKeyName,omitempty" xml:"DataKeyName,omitempty"`
	// The version number of the returned managed data key (DK).
	//
	// example:
	//
	// xH6OPUmz
	DataKeyVersionId *string `json:"DataKeyVersionId,omitempty" xml:"DataKeyVersionId,omitempty"`
	// The credential name that stores the key material of the returned managed data key (DK) version.
	//
	// example:
	//
	// kms-datakeyversion!example-data-key!xH6OPUmz
	DataKeyVersionName *string `json:"DataKeyVersionName,omitempty" xml:"DataKeyVersionName,omitempty"`
	// The Base64-encoding plaintext value of the data key (DK).
	//
	// example:
	//
	// CYueyVmZJ2MfA1VSZV2jCbFT8bO7StAvBnHacplr9aI=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 4bd560a1-729e-45f1-a3d9-b2a33d61046b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetManagedDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedDataKeyResponseBody) GetDataKeyName() *string {
	return s.DataKeyName
}

func (s *GetManagedDataKeyResponseBody) GetDataKeyVersionId() *string {
	return s.DataKeyVersionId
}

func (s *GetManagedDataKeyResponseBody) GetDataKeyVersionName() *string {
	return s.DataKeyVersionName
}

func (s *GetManagedDataKeyResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *GetManagedDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedDataKeyResponseBody) SetDataKeyName(v string) *GetManagedDataKeyResponseBody {
	s.DataKeyName = &v
	return s
}

func (s *GetManagedDataKeyResponseBody) SetDataKeyVersionId(v string) *GetManagedDataKeyResponseBody {
	s.DataKeyVersionId = &v
	return s
}

func (s *GetManagedDataKeyResponseBody) SetDataKeyVersionName(v string) *GetManagedDataKeyResponseBody {
	s.DataKeyVersionName = &v
	return s
}

func (s *GetManagedDataKeyResponseBody) SetPlaintext(v string) *GetManagedDataKeyResponseBody {
	s.Plaintext = &v
	return s
}

func (s *GetManagedDataKeyResponseBody) SetRequestId(v string) *GetManagedDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
