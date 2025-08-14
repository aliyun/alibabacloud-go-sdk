// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleBatchOssPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetHost() *string
	SetKey(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeSampleBatchOssPolicyResponseBody
	GetSignature() *string
}

type DescribeSampleBatchOssPolicyResponseBody struct {
	// OSS Access ID
	//
	// example:
	//
	// LTAxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// OSS Domain
	//
	// example:
	//
	// 172.25.126.234
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Key required for file upload.
	//
	// example:
	//
	// saf/de/namelist/e924/ufzgsedX9bd3a7
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// OSS Security Policy
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wNy0zMFQwNjowNTo0OS45NTRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJlcSIsIiRrZXkiLCJzYWZcL2RlXC9uYW1lbGlzdFwvZTkyNFwvdWZ6Z3NlZFg5Ymxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// OSS Signature
	//
	// example:
	//
	// PoAUQ//RusJJFIvCrn36O3+mM/U=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeSampleBatchOssPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleBatchOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleBatchOssPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetAccessId(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetHost(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetKey(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetPolicy(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetRequestId(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) SetSignature(v string) *DescribeSampleBatchOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
