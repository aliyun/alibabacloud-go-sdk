// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOssPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeModelOssPolicyResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeModelOssPolicyResponseBody
	GetHost() *string
	SetKey(v string) *DescribeModelOssPolicyResponseBody
	GetKey() *string
	SetMessage(v string) *DescribeModelOssPolicyResponseBody
	GetMessage() *string
	SetPolicy(v string) *DescribeModelOssPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeModelOssPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeModelOssPolicyResponseBody
	GetSignature() *string
	SetResultObject(v bool) *DescribeModelOssPolicyResponseBody
	GetResultObject() *bool
}

type DescribeModelOssPolicyResponseBody struct {
	// Access ID for OSS.
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// Address.
	//
	// example:
	//
	// https://xxxxxxxx-xxxxxxx.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// OSS access key secret.
	//
	// example:
	//
	// saf/a/uid/ccc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The policy for user form upload, which is a base64-encoded string.
	//
	// example:
	//
	// eyJleHBpxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Signature information.
	//
	// example:
	//
	// tzl1wL4q8rR/xxxxxx
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeModelOssPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelOssPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeModelOssPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeModelOssPolicyResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeModelOssPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeModelOssPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeModelOssPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelOssPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeModelOssPolicyResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeModelOssPolicyResponseBody) SetAccessId(v string) *DescribeModelOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetHost(v string) *DescribeModelOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetKey(v string) *DescribeModelOssPolicyResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetMessage(v string) *DescribeModelOssPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetPolicy(v string) *DescribeModelOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetRequestId(v string) *DescribeModelOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetSignature(v string) *DescribeModelOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) SetResultObject(v bool) *DescribeModelOssPolicyResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeModelOssPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
