// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyIds(v *DescribeEncryptionKeyListResponseBodyKeyIds) *DescribeEncryptionKeyListResponseBody
	GetKeyIds() *DescribeEncryptionKeyListResponseBodyKeyIds
	SetRequestId(v string) *DescribeEncryptionKeyListResponseBody
	GetRequestId() *string
}

type DescribeEncryptionKeyListResponseBody struct {
	// The custom keys that are available in the region.
	KeyIds *DescribeEncryptionKeyListResponseBodyKeyIds `json:"KeyIds,omitempty" xml:"KeyIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 10E2160B-959C-5C3E-BFE6-86EC5925****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyListResponseBody) GetKeyIds() *DescribeEncryptionKeyListResponseBodyKeyIds {
	return s.KeyIds
}

func (s *DescribeEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEncryptionKeyListResponseBody) SetKeyIds(v *DescribeEncryptionKeyListResponseBodyKeyIds) *DescribeEncryptionKeyListResponseBody {
	s.KeyIds = v
	return s
}

func (s *DescribeEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEncryptionKeyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEncryptionKeyListResponseBodyKeyIds struct {
	KeyId []*string `json:"KeyId,omitempty" xml:"KeyId,omitempty" type:"Repeated"`
}

func (s DescribeEncryptionKeyListResponseBodyKeyIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyListResponseBodyKeyIds) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyListResponseBodyKeyIds) GetKeyId() []*string {
	return s.KeyId
}

func (s *DescribeEncryptionKeyListResponseBodyKeyIds) SetKeyId(v []*string) *DescribeEncryptionKeyListResponseBodyKeyIds {
	s.KeyId = v
	return s
}

func (s *DescribeEncryptionKeyListResponseBodyKeyIds) Validate() error {
	return dara.Validate(s)
}
