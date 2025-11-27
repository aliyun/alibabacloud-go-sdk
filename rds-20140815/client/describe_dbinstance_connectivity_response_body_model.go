// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnCheckErrorCode(v string) *DescribeDBInstanceConnectivityResponseBody
	GetConnCheckErrorCode() *string
	SetConnCheckErrorMessage(v string) *DescribeDBInstanceConnectivityResponseBody
	GetConnCheckErrorMessage() *string
	SetConnCheckResult(v string) *DescribeDBInstanceConnectivityResponseBody
	GetConnCheckResult() *string
	SetDbInstanceName(v string) *DescribeDBInstanceConnectivityResponseBody
	GetDbInstanceName() *string
	SetRequestId(v string) *DescribeDBInstanceConnectivityResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceConnectivityResponseBody struct {
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
	// Src ip:39.106.64.59 not in user whitelist
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
	// The instance ID.
	//
	// example:
	//
	// rm-2ze2za3is7baay1w4
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D880212A-F21F-5722-8422-BD06B2874CC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConnectivityResponseBody) GetConnCheckErrorCode() *string {
	return s.ConnCheckErrorCode
}

func (s *DescribeDBInstanceConnectivityResponseBody) GetConnCheckErrorMessage() *string {
	return s.ConnCheckErrorMessage
}

func (s *DescribeDBInstanceConnectivityResponseBody) GetConnCheckResult() *string {
	return s.ConnCheckResult
}

func (s *DescribeDBInstanceConnectivityResponseBody) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstanceConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceConnectivityResponseBody) SetConnCheckErrorCode(v string) *DescribeDBInstanceConnectivityResponseBody {
	s.ConnCheckErrorCode = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponseBody) SetConnCheckErrorMessage(v string) *DescribeDBInstanceConnectivityResponseBody {
	s.ConnCheckErrorMessage = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponseBody) SetConnCheckResult(v string) *DescribeDBInstanceConnectivityResponseBody {
	s.ConnCheckResult = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponseBody) SetDbInstanceName(v string) *DescribeDBInstanceConnectivityResponseBody {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponseBody) SetRequestId(v string) *DescribeDBInstanceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
