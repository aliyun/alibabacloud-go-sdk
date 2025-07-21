// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReEncryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *ReEncryptResponseBody
	GetCiphertextBlob() *string
	SetKeyId(v string) *ReEncryptResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *ReEncryptResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *ReEncryptResponseBody
	GetRequestId() *string
}

type ReEncryptResponseBody struct {
	// The ciphertext re-encrypted.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK that is used to decrypt the original ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the original ciphertext.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 207596a2-36d3-4840-b1bd-f87044699bd7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReEncryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *ReEncryptResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *ReEncryptResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *ReEncryptResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *ReEncryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReEncryptResponseBody) SetCiphertextBlob(v string) *ReEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *ReEncryptResponseBody) SetKeyId(v string) *ReEncryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *ReEncryptResponseBody) SetKeyVersionId(v string) *ReEncryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *ReEncryptResponseBody) SetRequestId(v string) *ReEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReEncryptResponseBody) Validate() error {
	return dara.Validate(s)
}
