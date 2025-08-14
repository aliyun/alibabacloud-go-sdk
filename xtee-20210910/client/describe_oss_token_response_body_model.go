// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeOssTokenResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeOssTokenResponseBody
	GetHost() *string
	SetKey(v string) *DescribeOssTokenResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeOssTokenResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeOssTokenResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeOssTokenResponseBody
	GetSignature() *string
	SetStsToken(v string) *DescribeOssTokenResponseBody
	GetStsToken() *string
}

type DescribeOssTokenResponseBody struct {
	// AccessKeyId for OSS file upload.
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// Host address.
	//
	// example:
	//
	// kf.sunwoosoft.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The Key required for file upload.
	//
	// example:
	//
	// saf/de/namelist/e924/ufzgsedX9bd3a7
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Permission policy for ossbucket.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wNy0zMFQwNjowNTo0OS45NTRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJlcSIsIiRrZXkiLCJzYWZcL2RlXC9uYW1lbGlzdFwvZTkyNFwvdWZ6Z3NlZFg5Ymxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// OSS signature.
	//
	// example:
	//
	// n29by5MWBmAjcweVoPEY/OHktog=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Temporary identity credential.
	//
	// example:
	//
	// tT44bMQxxxxxxxxxxxxxxx
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s DescribeOssTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssTokenResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeOssTokenResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeOssTokenResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeOssTokenResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeOssTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssTokenResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeOssTokenResponseBody) GetStsToken() *string {
	return s.StsToken
}

func (s *DescribeOssTokenResponseBody) SetAccessId(v string) *DescribeOssTokenResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetHost(v string) *DescribeOssTokenResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetKey(v string) *DescribeOssTokenResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetPolicy(v string) *DescribeOssTokenResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetRequestId(v string) *DescribeOssTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetSignature(v string) *DescribeOssTokenResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeOssTokenResponseBody) SetStsToken(v string) *DescribeOssTokenResponseBody {
	s.StsToken = &v
	return s
}

func (s *DescribeOssTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
