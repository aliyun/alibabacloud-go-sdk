// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeEventUploadPolicyResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeEventUploadPolicyResponseBody
	GetHost() *string
	SetKey(v string) *DescribeEventUploadPolicyResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeEventUploadPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeEventUploadPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeEventUploadPolicyResponseBody
	GetSignature() *string
	SetStsToken(v string) *DescribeEventUploadPolicyResponseBody
	GetStsToken() *string
}

type DescribeEventUploadPolicyResponseBody struct {
	// ID for accessing OSS
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// OSS host.
	//
	// example:
	//
	// 172.16.0.44
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The Key required for file upload.
	//
	// example:
	//
	// saf/de/namelist/e924/ufzgsedX9bd3a7
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// OSS security policy
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wNy0zMFQwNjowNTo0OS45NTRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJlcSIsIiRrZXkiLCJzYWZcL2RlXC9uYW1lbGlzdFwvZTkyNFwvdWZ6Z3NlZFg5Ymxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Signature data.
	//
	// example:
	//
	// 7aXmkd2Z3oksCXOS9uvKlJuOKaY=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Temporary identity credential.
	//
	// example:
	//
	// tT44bMQxxxxxxxxxxxxxxx
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s DescribeEventUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventUploadPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeEventUploadPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeEventUploadPolicyResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeEventUploadPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeEventUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventUploadPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeEventUploadPolicyResponseBody) GetStsToken() *string {
	return s.StsToken
}

func (s *DescribeEventUploadPolicyResponseBody) SetAccessId(v string) *DescribeEventUploadPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetHost(v string) *DescribeEventUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetKey(v string) *DescribeEventUploadPolicyResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetPolicy(v string) *DescribeEventUploadPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetRequestId(v string) *DescribeEventUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetSignature(v string) *DescribeEventUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) SetStsToken(v string) *DescribeEventUploadPolicyResponseBody {
	s.StsToken = &v
	return s
}

func (s *DescribeEventUploadPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
