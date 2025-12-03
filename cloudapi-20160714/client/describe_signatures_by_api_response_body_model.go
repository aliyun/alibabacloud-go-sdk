// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesByApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSignaturesByApiResponseBody
	GetRequestId() *string
	SetSignatures(v *DescribeSignaturesByApiResponseBodySignatures) *DescribeSignaturesByApiResponseBody
	GetSignatures() *DescribeSignaturesByApiResponseBodySignatures
}

type DescribeSignaturesByApiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned signature key information. It is an array consisting of SignatureItem data.
	Signatures *DescribeSignaturesByApiResponseBodySignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Struct"`
}

func (s DescribeSignaturesByApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSignaturesByApiResponseBody) GetSignatures() *DescribeSignaturesByApiResponseBodySignatures {
	return s.Signatures
}

func (s *DescribeSignaturesByApiResponseBody) SetRequestId(v string) *DescribeSignaturesByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBody) SetSignatures(v *DescribeSignaturesByApiResponseBodySignatures) *DescribeSignaturesByApiResponseBody {
	s.Signatures = v
	return s
}

func (s *DescribeSignaturesByApiResponseBody) Validate() error {
	if s.Signatures != nil {
		if err := s.Signatures.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSignaturesByApiResponseBodySignatures struct {
	SignatureItem []*DescribeSignaturesByApiResponseBodySignaturesSignatureItem `json:"SignatureItem,omitempty" xml:"SignatureItem,omitempty" type:"Repeated"`
}

func (s DescribeSignaturesByApiResponseBodySignatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBodySignatures) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBodySignatures) GetSignatureItem() []*DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	return s.SignatureItem
}

func (s *DescribeSignaturesByApiResponseBodySignatures) SetSignatureItem(v []*DescribeSignaturesByApiResponseBodySignaturesSignatureItem) *DescribeSignaturesByApiResponseBodySignatures {
	s.SignatureItem = v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignatures) Validate() error {
	if s.SignatureItem != nil {
		for _, item := range s.SignatureItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSignaturesByApiResponseBodySignaturesSignatureItem struct {
	// The time when the key was bound.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	BoundTime *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	// The ID of the backend signature key.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The name of the backend signature key.
	//
	// example:
	//
	// mysecret
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeSignaturesByApiResponseBodySignaturesSignatureItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBodySignaturesSignatureItem) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) GetSignatureName() *string {
	return s.SignatureName
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetBoundTime(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetSignatureId(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetSignatureName(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.SignatureName = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) Validate() error {
	return dara.Validate(s)
}
