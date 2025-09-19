// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *DescribeMonitorAccountsResponseBody
	GetAccountIds() []*string
	SetRequestId(v string) *DescribeMonitorAccountsResponseBody
	GetRequestId() *string
}

type DescribeMonitorAccountsResponseBody struct {
	// The IDs of the members.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The request ID.
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

func (s *DescribeMonitorAccountsResponseBody) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *DescribeMonitorAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
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
	return dara.Validate(s)
}
