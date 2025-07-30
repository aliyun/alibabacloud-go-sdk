// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificatePrivateKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedData(v string) *DescribeCertificatePrivateKeyResponseBody
	GetEncryptedData() *string
	SetRequestId(v string) *DescribeCertificatePrivateKeyResponseBody
	GetRequestId() *string
}

type DescribeCertificatePrivateKeyResponseBody struct {
	// The content of the encrypted private key.
	//
	// example:
	//
	// -----BEGIN ENCRYPTED PRIVATE KEY----- …… -----END ENCRYPTED PRIVATE KEY-----
	EncryptedData *string `json:"EncryptedData,omitempty" xml:"EncryptedData,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertificatePrivateKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificatePrivateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificatePrivateKeyResponseBody) GetEncryptedData() *string {
	return s.EncryptedData
}

func (s *DescribeCertificatePrivateKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertificatePrivateKeyResponseBody) SetEncryptedData(v string) *DescribeCertificatePrivateKeyResponseBody {
	s.EncryptedData = &v
	return s
}

func (s *DescribeCertificatePrivateKeyResponseBody) SetRequestId(v string) *DescribeCertificatePrivateKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificatePrivateKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
