// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainBackupResponseBody
	GetDomainName() *string
	SetPeriodType(v string) *AddDomainBackupResponseBody
	GetPeriodType() *string
	SetRequestId(v string) *AddDomainBackupResponseBody
	GetRequestId() *string
}

type AddDomainBackupResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// test.aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The backup cycle.
	//
	// example:
	//
	// DAY
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FD552816-FCC8-4832-B4A2-2DA0C2BA1688
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDomainBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainBackupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainBackupResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainBackupResponseBody) GetPeriodType() *string {
	return s.PeriodType
}

func (s *AddDomainBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainBackupResponseBody) SetDomainName(v string) *AddDomainBackupResponseBody {
	s.DomainName = &v
	return s
}

func (s *AddDomainBackupResponseBody) SetPeriodType(v string) *AddDomainBackupResponseBody {
	s.PeriodType = &v
	return s
}

func (s *AddDomainBackupResponseBody) SetRequestId(v string) *AddDomainBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
