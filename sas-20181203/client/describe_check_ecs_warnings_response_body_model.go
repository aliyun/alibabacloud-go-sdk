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
	// Indicates whether the current Security Center edition is a trial version. Valid values:
	//
	// - **0**: not a trial version
	//
	// - **1**: a trial version.
	//
	// example:
	//
	// 0
	CanTry *string `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The edition of Security Center that you purchased. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **2*	- or **3**: Enterprise Edition
	//
	// - **5**: Premium Edition
	//
	// - **6**: Anti-virus Edition
	//
	// > Both 2 and 3 correspond to Enterprise Edition. There is no difference between the two values.
	//
	// example:
	//
	// 3
	SasVersion *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
	// The number of high-risk weak password risks detected in your assets.
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
