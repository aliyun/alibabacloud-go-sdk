// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeSampleUploadPolicyResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeSampleUploadPolicyResponseBody
	GetHost() *string
	SetKey(v string) *DescribeSampleUploadPolicyResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeSampleUploadPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeSampleUploadPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeSampleUploadPolicyResponseBody
	GetSignature() *string
	SetStsToken(v string) *DescribeSampleUploadPolicyResponseBody
	GetStsToken() *string
}

type DescribeSampleUploadPolicyResponseBody struct {
	// OSS access key ID.
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// OSS domain name.
	//
	// example:
	//
	// emseu.cxy8uoq4aafx.eu-central-1.rds.amazonaws.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The key of OSS upload policy.
	//
	// example:
	//
	// saf/de/namelist/e924/ufzgsedX9bd3a7
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// OSS upload policy.
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
	// Signature information.
	//
	// example:
	//
	// 3Es5j/9Xm/zwPcM9cwEr5pa0Wsc=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Temporary identity credential.
	//
	// example:
	//
	// tT44bMQxxxxxxxxxxxxxxx
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s DescribeSampleUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleUploadPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeSampleUploadPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeSampleUploadPolicyResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeSampleUploadPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeSampleUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleUploadPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeSampleUploadPolicyResponseBody) GetStsToken() *string {
	return s.StsToken
}

func (s *DescribeSampleUploadPolicyResponseBody) SetAccessId(v string) *DescribeSampleUploadPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetHost(v string) *DescribeSampleUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetKey(v string) *DescribeSampleUploadPolicyResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetPolicy(v string) *DescribeSampleUploadPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetRequestId(v string) *DescribeSampleUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetSignature(v string) *DescribeSampleUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) SetStsToken(v string) *DescribeSampleUploadPolicyResponseBody {
	s.StsToken = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
