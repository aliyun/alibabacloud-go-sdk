// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdInfos(v []*DescribeMonitorAccountsResponseBodyAccountIdInfos) *DescribeMonitorAccountsResponseBody
	GetAccountIdInfos() []*DescribeMonitorAccountsResponseBodyAccountIdInfos
	SetAccountIds(v []*string) *DescribeMonitorAccountsResponseBody
	GetAccountIds() []*string
	SetRequestId(v string) *DescribeMonitorAccountsResponseBody
	GetRequestId() *string
}

type DescribeMonitorAccountsResponseBody struct {
	// The list of member account information.
	AccountIdInfos []*DescribeMonitorAccountsResponseBodyAccountIdInfos `json:"AccountIdInfos,omitempty" xml:"AccountIdInfos,omitempty" type:"Repeated"`
	// The list of member account IDs.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BF880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMonitorAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorAccountsResponseBody) GetAccountIdInfos() []*DescribeMonitorAccountsResponseBodyAccountIdInfos {
	return s.AccountIdInfos
}

func (s *DescribeMonitorAccountsResponseBody) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *DescribeMonitorAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorAccountsResponseBody) SetAccountIdInfos(v []*DescribeMonitorAccountsResponseBodyAccountIdInfos) *DescribeMonitorAccountsResponseBody {
	s.AccountIdInfos = v
	return s
}

func (s *DescribeMonitorAccountsResponseBody) SetAccountIds(v []*string) *DescribeMonitorAccountsResponseBody {
	s.AccountIds = v
	return s
}

func (s *DescribeMonitorAccountsResponseBody) SetRequestId(v string) *DescribeMonitorAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBody) Validate() error {
	if s.AccountIdInfos != nil {
		for _, item := range s.AccountIdInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorAccountsResponseBodyAccountIdInfos struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 119593010538****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the account was added to the management list. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1760520684000
	AddTime *int64 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The Alibaba Cloud account ID of the administrator.
	//
	// example:
	//
	// 106635707417****
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The basic service switch. Valid values:
	//
	// - **0**: off
	//
	// - **1**: on
	//
	// example:
	//
	// 1
	PostBasicService *int32 `json:"PostBasicService,omitempty" xml:"PostBasicService,omitempty"`
	// The purchased edition of Security Center. Valid values:
	//
	// - **0*	- or **1**: Free Edition
	//
	// - **2*	- or **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// example:
	//
	// 7
	SasVersion *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
}

func (s DescribeMonitorAccountsResponseBodyAccountIdInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorAccountsResponseBodyAccountIdInfos) GoString() string {
	return s.String()
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) GetAddTime() *int64 {
	return s.AddTime
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) GetPostBasicService() *int32 {
	return s.PostBasicService
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) GetSasVersion() *string {
	return s.SasVersion
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) SetAccountId(v string) *DescribeMonitorAccountsResponseBodyAccountIdInfos {
	s.AccountId = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) SetAddTime(v int64) *DescribeMonitorAccountsResponseBodyAccountIdInfos {
	s.AddTime = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) SetOperatorUid(v string) *DescribeMonitorAccountsResponseBodyAccountIdInfos {
	s.OperatorUid = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) SetPostBasicService(v int32) *DescribeMonitorAccountsResponseBodyAccountIdInfos {
	s.PostBasicService = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) SetSasVersion(v string) *DescribeMonitorAccountsResponseBodyAccountIdInfos {
	s.SasVersion = &v
	return s
}

func (s *DescribeMonitorAccountsResponseBodyAccountIdInfos) Validate() error {
	return dara.Validate(s)
}
