// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckEcsWarningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanTry(v string) *DescribeCheckEcsWarningsResponseBody
	GetCanTry() *string
	SetRequestId(v string) *DescribeCheckEcsWarningsResponseBody
	GetRequestId() *string
	SetSasVersion(v string) *DescribeCheckEcsWarningsResponseBody
	GetSasVersion() *string
	SetWeakPasswordCount(v string) *DescribeCheckEcsWarningsResponseBody
	GetWeakPasswordCount() *string
}

type DescribeCheckEcsWarningsResponseBody struct {
	// Indicates whether you use the free trial of Security Center. Valid values:
	//
	// - **0**: no
	//
	// - **1**: yes
	//
	// example:
	//
	// 0
	CanTry *string `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The edition of Security Center that you use. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **2*	- or **3**: Enterprise edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **6**: Anti-virus edition
	//
	// >  Both the value 2 and the value 3 indicate the Enterprise edition.
	//
	// example:
	//
	// 3
	SasVersion *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
	// The number of weak passwords that can cause high risks to your assets.
	//
	// example:
	//
	// 3
	WeakPasswordCount *string `json:"WeakPasswordCount,omitempty" xml:"WeakPasswordCount,omitempty"`
}

func (s DescribeCheckEcsWarningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckEcsWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsResponseBody) GetCanTry() *string {
	return s.CanTry
}

func (s *DescribeCheckEcsWarningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckEcsWarningsResponseBody) GetSasVersion() *string {
	return s.SasVersion
}

func (s *DescribeCheckEcsWarningsResponseBody) GetWeakPasswordCount() *string {
	return s.WeakPasswordCount
}

func (s *DescribeCheckEcsWarningsResponseBody) SetCanTry(v string) *DescribeCheckEcsWarningsResponseBody {
	s.CanTry = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetRequestId(v string) *DescribeCheckEcsWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetSasVersion(v string) *DescribeCheckEcsWarningsResponseBody {
	s.SasVersion = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetWeakPasswordCount(v string) *DescribeCheckEcsWarningsResponseBody {
	s.WeakPasswordCount = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) Validate() error {
	return dara.Validate(s)
}
