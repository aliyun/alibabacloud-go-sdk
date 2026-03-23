// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSQLCollectorPolicyResponseBody
	GetRequestId() *string
	SetSQLCollectorStatus(v string) *DescribeSQLCollectorPolicyResponseBody
	GetSQLCollectorStatus() *string
	SetStoragePeriod(v int32) *DescribeSQLCollectorPolicyResponseBody
	GetStoragePeriod() *int32
}

type DescribeSQLCollectorPolicyResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
	StoragePeriod      *int32  `json:"StoragePeriod,omitempty" xml:"StoragePeriod,omitempty"`
}

func (s DescribeSQLCollectorPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLCollectorPolicyResponseBody) GetSQLCollectorStatus() *string {
	return s.SQLCollectorStatus
}

func (s *DescribeSQLCollectorPolicyResponseBody) GetStoragePeriod() *int32 {
	return s.StoragePeriod
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetRequestId(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribeSQLCollectorPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponseBody) SetStoragePeriod(v int32) *DescribeSQLCollectorPolicyResponseBody {
	s.StoragePeriod = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
