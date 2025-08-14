// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeOssPolicyResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeOssPolicyResponseBody
	GetHost() *string
	SetKey(v string) *DescribeOssPolicyResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeOssPolicyResponseBody
	GetPolicy() *string
	SetSignature(v string) *DescribeOssPolicyResponseBody
	GetSignature() *string
	SetRequestId(v string) *DescribeOssPolicyResponseBody
	GetRequestId() *string
}

type DescribeOssPolicyResponseBody struct {
	// accessId, a parameter used in OSS SDK uploads, corresponding to OSSAccessKeyId
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// OSS host ID.
	//
	// example:
	//
	// testvm.biubiubiuu.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Key required for file upload.
	//
	// example:
	//
	// saf/a/uid/ccc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// OSS security policy
	//
	// example:
	//
	// eyJleHBpxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Signature data.
	//
	// example:
	//
	// tzl1wL4q8rR/xxxxxx
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeOssPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeOssPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeOssPolicyResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeOssPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeOssPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeOssPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssPolicyResponseBody) SetAccessId(v string) *DescribeOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) SetHost(v string) *DescribeOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) SetKey(v string) *DescribeOssPolicyResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) SetPolicy(v string) *DescribeOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) SetSignature(v string) *DescribeOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) SetRequestId(v string) *DescribeOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
