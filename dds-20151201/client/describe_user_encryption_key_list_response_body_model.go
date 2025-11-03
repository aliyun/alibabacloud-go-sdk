// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyIds(v *DescribeUserEncryptionKeyListResponseBodyKeyIds) *DescribeUserEncryptionKeyListResponseBody
	GetKeyIds() *DescribeUserEncryptionKeyListResponseBodyKeyIds
	SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetRequestId() *string
}

type DescribeUserEncryptionKeyListResponseBody struct {
	// The list of custom keys.
	KeyIds *DescribeUserEncryptionKeyListResponseBodyKeyIds `json:"KeyIds,omitempty" xml:"KeyIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7CD51DA4-A499-43CE-B9B5-20CD4FDC648E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetKeyIds() *DescribeUserEncryptionKeyListResponseBodyKeyIds {
	return s.KeyIds
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKeyIds(v *DescribeUserEncryptionKeyListResponseBodyKeyIds) *DescribeUserEncryptionKeyListResponseBody {
	s.KeyIds = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) Validate() error {
	if s.KeyIds != nil {
		if err := s.KeyIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserEncryptionKeyListResponseBodyKeyIds struct {
	KeyId []*string `json:"KeyId,omitempty" xml:"KeyId,omitempty" type:"Repeated"`
}

func (s DescribeUserEncryptionKeyListResponseBodyKeyIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyKeyIds) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyKeyIds) GetKeyId() []*string {
	return s.KeyId
}

func (s *DescribeUserEncryptionKeyListResponseBodyKeyIds) SetKeyId(v []*string) *DescribeUserEncryptionKeyListResponseBodyKeyIds {
	s.KeyId = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBodyKeyIds) Validate() error {
	return dara.Validate(s)
}
