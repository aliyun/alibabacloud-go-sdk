// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSQLExplorerPolicyResponseBody
	GetRequestId() *string
	SetSQLCollectorStatus(v string) *DescribeSQLExplorerPolicyResponseBody
	GetSQLCollectorStatus() *string
}

type DescribeSQLExplorerPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3655211B-4D74-4E13-91E6-FF2AFE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the SQL Explorer feature is enabled. Valid values:
	//
	// 	- **Enable**: Enabled.
	//
	// 	- **Disabled**: Disabled.
	//
	// example:
	//
	// Enable
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s DescribeSQLExplorerPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLExplorerPolicyResponseBody) GetSQLCollectorStatus() *string {
	return s.SQLCollectorStatus
}

func (s *DescribeSQLExplorerPolicyResponseBody) SetRequestId(v string) *DescribeSQLExplorerPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLExplorerPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribeSQLExplorerPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

func (s *DescribeSQLExplorerPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
