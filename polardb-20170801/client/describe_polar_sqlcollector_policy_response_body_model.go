// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarSQLCollectorPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarSQLCollectorPolicyResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribePolarSQLCollectorPolicyResponseBody
	GetRequestId() *string
	SetSQLCollectorStatus(v string) *DescribePolarSQLCollectorPolicyResponseBody
	GetSQLCollectorStatus() *string
}

type DescribePolarSQLCollectorPolicyResponseBody struct {
	// The IDs of the clusters.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3655211B-4D74-4E13-91E6-FF2AFE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the SQL Explorer feature is enabled. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enable
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s DescribePolarSQLCollectorPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) GetSQLCollectorStatus() *string {
	return s.SQLCollectorStatus
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) SetDBClusterId(v string) *DescribePolarSQLCollectorPolicyResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) SetRequestId(v string) *DescribePolarSQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribePolarSQLCollectorPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
