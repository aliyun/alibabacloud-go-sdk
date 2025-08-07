// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnCheckErrorCode(v string) *DescribeDBClusterConnectivityResponseBody
	GetConnCheckErrorCode() *string
	SetConnCheckErrorMessage(v string) *DescribeDBClusterConnectivityResponseBody
	GetConnCheckErrorMessage() *string
	SetConnCheckResult(v string) *DescribeDBClusterConnectivityResponseBody
	GetConnCheckResult() *string
	SetDBClusterId(v string) *DescribeDBClusterConnectivityResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeDBClusterConnectivityResponseBody
	GetRequestId() *string
}

type DescribeDBClusterConnectivityResponseBody struct {
	// The error code for connection diagnosis. Valid values:
	//
	// 	- **SRC_IP_NOT_IN_USER_WHITELIST**: The source IP address is not added to the whitelist.
	//
	// 	- **CONNECTION_ABNORMAL**: The connection to the cluster is normal.
	//
	// example:
	//
	// SRC_IP_NOT_IN_USER_WHITELIST
	ConnCheckErrorCode *string `json:"ConnCheckErrorCode,omitempty" xml:"ConnCheckErrorCode,omitempty"`
	// The error message for connection diagnosis.
	//
	// example:
	//
	// Src ip:192.***.***.1 not in user whitelist
	ConnCheckErrorMessage *string `json:"ConnCheckErrorMessage,omitempty" xml:"ConnCheckErrorMessage,omitempty"`
	// The connection diagnosis result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Failed
	ConnCheckResult *string `json:"ConnCheckResult,omitempty" xml:"ConnCheckResult,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-xxxxxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73A85BAF-1039-4CDE-A83F-1A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConnectivityResponseBody) GetConnCheckErrorCode() *string {
	return s.ConnCheckErrorCode
}

func (s *DescribeDBClusterConnectivityResponseBody) GetConnCheckErrorMessage() *string {
	return s.ConnCheckErrorMessage
}

func (s *DescribeDBClusterConnectivityResponseBody) GetConnCheckResult() *string {
	return s.ConnCheckResult
}

func (s *DescribeDBClusterConnectivityResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterConnectivityResponseBody) SetConnCheckErrorCode(v string) *DescribeDBClusterConnectivityResponseBody {
	s.ConnCheckErrorCode = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponseBody) SetConnCheckErrorMessage(v string) *DescribeDBClusterConnectivityResponseBody {
	s.ConnCheckErrorMessage = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponseBody) SetConnCheckResult(v string) *DescribeDBClusterConnectivityResponseBody {
	s.ConnCheckResult = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponseBody) SetDBClusterId(v string) *DescribeDBClusterConnectivityResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponseBody) SetRequestId(v string) *DescribeDBClusterConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
