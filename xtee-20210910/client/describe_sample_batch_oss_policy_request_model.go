// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleBatchOssPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleBatchOssPolicyRequest
	GetLang() *string
	SetBatchName(v string) *DescribeSampleBatchOssPolicyRequest
	GetBatchName() *string
	SetRegId(v string) *DescribeSampleBatchOssPolicyRequest
	GetRegId() *string
}

type DescribeSampleBatchOssPolicyRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Sample batch name
	//
	// example:
	//
	// 白样本
	BatchName *string `json:"batchName,omitempty" xml:"batchName,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSampleBatchOssPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleBatchOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleBatchOssPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleBatchOssPolicyRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *DescribeSampleBatchOssPolicyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleBatchOssPolicyRequest) SetLang(v string) *DescribeSampleBatchOssPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyRequest) SetBatchName(v string) *DescribeSampleBatchOssPolicyRequest {
	s.BatchName = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyRequest) SetRegId(v string) *DescribeSampleBatchOssPolicyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyRequest) Validate() error {
	return dara.Validate(s)
}
