// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFingerprint(v string) *DescribeInstanceKeyPairResponseBody
	GetFingerprint() *string
	SetKeyPairName(v string) *DescribeInstanceKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *DescribeInstanceKeyPairResponseBody
	GetRequestId() *string
}

type DescribeInstanceKeyPairResponseBody struct {
	// The fingerprint of the key pair.
	//
	// example:
	//
	// 4f:70:62:e9:0c:72:f7:ee:74:ce:e3:bf:e0:82:**:**
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeInstanceKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeInstanceKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceKeyPairResponseBody) SetFingerprint(v string) *DescribeInstanceKeyPairResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *DescribeInstanceKeyPairResponseBody) SetKeyPairName(v string) *DescribeInstanceKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DescribeInstanceKeyPairResponseBody) SetRequestId(v string) *DescribeInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
