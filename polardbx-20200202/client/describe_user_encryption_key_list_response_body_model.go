// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserEncryptionKeyListResponseBodyData) *DescribeUserEncryptionKeyListResponseBody
	GetData() *DescribeUserEncryptionKeyListResponseBodyData
	SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetRequestId() *string
}

type DescribeUserEncryptionKeyListResponseBody struct {
	Data *DescribeUserEncryptionKeyListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetData() *DescribeUserEncryptionKeyListResponseBodyData {
	return s.Data
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetData(v *DescribeUserEncryptionKeyListResponseBodyData) *DescribeUserEncryptionKeyListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserEncryptionKeyListResponseBodyData struct {
	KeyIds []*string `json:"KeyIds,omitempty" xml:"KeyIds,omitempty" type:"Repeated"`
}

func (s DescribeUserEncryptionKeyListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyData) GetKeyIds() []*string {
	return s.KeyIds
}

func (s *DescribeUserEncryptionKeyListResponseBodyData) SetKeyIds(v []*string) *DescribeUserEncryptionKeyListResponseBodyData {
	s.KeyIds = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
