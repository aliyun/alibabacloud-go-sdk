// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstances(v *ListKmsInstancesResponseBodyKmsInstances) *ListKmsInstancesResponseBody
	GetKmsInstances() *ListKmsInstancesResponseBodyKmsInstances
	SetPageNumber(v int64) *ListKmsInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListKmsInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListKmsInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListKmsInstancesResponseBody
	GetTotalCount() *int64
}

type ListKmsInstancesResponseBody struct {
	// A list of KMS instances.
	KmsInstances *ListKmsInstancesResponseBodyKmsInstances `json:"KmsInstances,omitempty" xml:"KmsInstances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// d3eca5c8-a856-4347-8eb6-e1898c3fda2e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of KMS instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKmsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKmsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBody) GetKmsInstances() *ListKmsInstancesResponseBodyKmsInstances {
	return s.KmsInstances
}

func (s *ListKmsInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListKmsInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListKmsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKmsInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKmsInstancesResponseBody) SetKmsInstances(v *ListKmsInstancesResponseBodyKmsInstances) *ListKmsInstancesResponseBody {
	s.KmsInstances = v
	return s
}

func (s *ListKmsInstancesResponseBody) SetPageNumber(v int64) *ListKmsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetPageSize(v int64) *ListKmsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetRequestId(v string) *ListKmsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKmsInstancesResponseBody) SetTotalCount(v int64) *ListKmsInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKmsInstancesResponseBody) Validate() error {
	if s.KmsInstances != nil {
		if err := s.KmsInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKmsInstancesResponseBodyKmsInstances struct {
	KmsInstance []*ListKmsInstancesResponseBodyKmsInstancesKmsInstance `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty" type:"Repeated"`
}

func (s ListKmsInstancesResponseBodyKmsInstances) String() string {
	return dara.Prettify(s)
}

func (s ListKmsInstancesResponseBodyKmsInstances) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBodyKmsInstances) GetKmsInstance() []*ListKmsInstancesResponseBodyKmsInstancesKmsInstance {
	return s.KmsInstance
}

func (s *ListKmsInstancesResponseBodyKmsInstances) SetKmsInstance(v []*ListKmsInstancesResponseBodyKmsInstancesKmsInstance) *ListKmsInstancesResponseBodyKmsInstances {
	s.KmsInstance = v
	return s
}

func (s *ListKmsInstancesResponseBodyKmsInstances) Validate() error {
	if s.KmsInstance != nil {
		for _, item := range s.KmsInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKmsInstancesResponseBodyKmsInstancesKmsInstance struct {
	// The ARN of the KMS instance.
	//
	// example:
	//
	// acs:kms:pre-hangzhou:120708975881****:keystore/kst-phzz64c9f84eo32dbs****
	KmsInstanceArn *string `json:"KmsInstanceArn,omitempty" xml:"KmsInstanceArn,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-phzz64c9f84eo32dbs****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s ListKmsInstancesResponseBodyKmsInstancesKmsInstance) String() string {
	return dara.Prettify(s)
}

func (s ListKmsInstancesResponseBodyKmsInstancesKmsInstance) GoString() string {
	return s.String()
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) GetKmsInstanceArn() *string {
	return s.KmsInstanceArn
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) SetKmsInstanceArn(v string) *ListKmsInstancesResponseBodyKmsInstancesKmsInstance {
	s.KmsInstanceArn = &v
	return s
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) SetKmsInstanceId(v string) *ListKmsInstancesResponseBodyKmsInstancesKmsInstance {
	s.KmsInstanceId = &v
	return s
}

func (s *ListKmsInstancesResponseBodyKmsInstancesKmsInstance) Validate() error {
	return dara.Validate(s)
}
