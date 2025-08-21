// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzerType(v string) *ListDictInformationRequest
	GetAnalyzerType() *string
	SetBucketName(v string) *ListDictInformationRequest
	GetBucketName() *string
	SetKey(v string) *ListDictInformationRequest
	GetKey() *string
}

type ListDictInformationRequest struct {
	// example:
	//
	// ALIWS
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// search-cloud-test-cn-****
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss/dic_0.dic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListDictInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDictInformationRequest) GoString() string {
	return s.String()
}

func (s *ListDictInformationRequest) GetAnalyzerType() *string {
	return s.AnalyzerType
}

func (s *ListDictInformationRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListDictInformationRequest) GetKey() *string {
	return s.Key
}

func (s *ListDictInformationRequest) SetAnalyzerType(v string) *ListDictInformationRequest {
	s.AnalyzerType = &v
	return s
}

func (s *ListDictInformationRequest) SetBucketName(v string) *ListDictInformationRequest {
	s.BucketName = &v
	return s
}

func (s *ListDictInformationRequest) SetKey(v string) *ListDictInformationRequest {
	s.Key = &v
	return s
}

func (s *ListDictInformationRequest) Validate() error {
	return dara.Validate(s)
}
