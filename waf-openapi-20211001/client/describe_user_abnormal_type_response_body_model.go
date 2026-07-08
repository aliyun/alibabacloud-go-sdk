// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormal(v []*DescribeUserAbnormalTypeResponseBodyAbnormal) *DescribeUserAbnormalTypeResponseBody
	GetAbnormal() []*DescribeUserAbnormalTypeResponseBodyAbnormal
	SetRequestId(v string) *DescribeUserAbnormalTypeResponseBody
	GetRequestId() *string
}

type DescribeUserAbnormalTypeResponseBody struct {
	// The risk types detected by API security and the number of risks for each type.
	Abnormal []*DescribeUserAbnormalTypeResponseBodyAbnormal `json:"Abnormal,omitempty" xml:"Abnormal,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3E1CB966-1407-5988-9432-7***D784
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserAbnormalTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTypeResponseBody) GetAbnormal() []*DescribeUserAbnormalTypeResponseBodyAbnormal {
	return s.Abnormal
}

func (s *DescribeUserAbnormalTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAbnormalTypeResponseBody) SetAbnormal(v []*DescribeUserAbnormalTypeResponseBodyAbnormal) *DescribeUserAbnormalTypeResponseBody {
	s.Abnormal = v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBody) SetRequestId(v string) *DescribeUserAbnormalTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBody) Validate() error {
	if s.Abnormal != nil {
		for _, item := range s.Abnormal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserAbnormalTypeResponseBodyAbnormal struct {
	// The code that identifies the specific risk, such as `Risk_InternalWeakPasswd`.
	//
	// example:
	//
	// Risk_InternalWeakPasswd
	AbnormalCode *string `json:"AbnormalCode,omitempty" xml:"AbnormalCode,omitempty"`
	// The total number of risks detected for this risk type.
	//
	// example:
	//
	// 10
	AbnormalCount *int64 `json:"AbnormalCount,omitempty" xml:"AbnormalCount,omitempty"`
	// The parent category of the risk, such as `RiskType_Account`.
	//
	// example:
	//
	// RiskType_Account
	AbnormalParentType *string `json:"AbnormalParentType,omitempty" xml:"AbnormalParentType,omitempty"`
	// The type of the risk detected by API security.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported risk types.
	//
	// example:
	//
	// Risk_InternalWeakPasswd
	AbnormalType *string `json:"AbnormalType,omitempty" xml:"AbnormalType,omitempty"`
}

func (s DescribeUserAbnormalTypeResponseBodyAbnormal) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTypeResponseBodyAbnormal) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) GetAbnormalCode() *string {
	return s.AbnormalCode
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) GetAbnormalCount() *int64 {
	return s.AbnormalCount
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) GetAbnormalParentType() *string {
	return s.AbnormalParentType
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) GetAbnormalType() *string {
	return s.AbnormalType
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) SetAbnormalCode(v string) *DescribeUserAbnormalTypeResponseBodyAbnormal {
	s.AbnormalCode = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) SetAbnormalCount(v int64) *DescribeUserAbnormalTypeResponseBodyAbnormal {
	s.AbnormalCount = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) SetAbnormalParentType(v string) *DescribeUserAbnormalTypeResponseBodyAbnormal {
	s.AbnormalParentType = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) SetAbnormalType(v string) *DescribeUserAbnormalTypeResponseBodyAbnormal {
	s.AbnormalType = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponseBodyAbnormal) Validate() error {
	return dara.Validate(s)
}
