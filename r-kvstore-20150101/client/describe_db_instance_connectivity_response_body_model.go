// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnCheckErrorCode(v string) *DescribeDbInstanceConnectivityResponseBody
	GetConnCheckErrorCode() *string
	SetConnCheckErrorMessage(v string) *DescribeDbInstanceConnectivityResponseBody
	GetConnCheckErrorMessage() *string
	SetConnCheckResult(v string) *DescribeDbInstanceConnectivityResponseBody
	GetConnCheckResult() *string
	SetRequestId(v string) *DescribeDbInstanceConnectivityResponseBody
	GetRequestId() *string
}

type DescribeDbInstanceConnectivityResponseBody struct {
	// The error code for connection diagnosis. Valid values:
	//
	// 	- **SRC_IP_NOT_IN_USER_WHITELIST**: The source IP address is not added to the whitelist.
	//
	// 	- **CONNECTION_ABNORMAL**: The connection to the instance is normal.
	//
	// example:
	//
	// SRC_IP_NOT_IN_USER_WHITELIST
	ConnCheckErrorCode *string `json:"ConnCheckErrorCode,omitempty" xml:"ConnCheckErrorCode,omitempty"`
	// The error message for connection diagnosis.
	//
	// example:
	//
	// Src ip:172.28.134.96 not in user whitelist
	ConnCheckErrorMessage *string `json:"ConnCheckErrorMessage,omitempty" xml:"ConnCheckErrorMessage,omitempty"`
	// The connection check result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Failed
	ConnCheckResult *string `json:"ConnCheckResult,omitempty" xml:"ConnCheckResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDbInstanceConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceConnectivityResponseBody) GetConnCheckErrorCode() *string {
	return s.ConnCheckErrorCode
}

func (s *DescribeDbInstanceConnectivityResponseBody) GetConnCheckErrorMessage() *string {
	return s.ConnCheckErrorMessage
}

func (s *DescribeDbInstanceConnectivityResponseBody) GetConnCheckResult() *string {
	return s.ConnCheckResult
}

func (s *DescribeDbInstanceConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDbInstanceConnectivityResponseBody) SetConnCheckErrorCode(v string) *DescribeDbInstanceConnectivityResponseBody {
	s.ConnCheckErrorCode = &v
	return s
}

func (s *DescribeDbInstanceConnectivityResponseBody) SetConnCheckErrorMessage(v string) *DescribeDbInstanceConnectivityResponseBody {
	s.ConnCheckErrorMessage = &v
	return s
}

func (s *DescribeDbInstanceConnectivityResponseBody) SetConnCheckResult(v string) *DescribeDbInstanceConnectivityResponseBody {
	s.ConnCheckResult = &v
	return s
}

func (s *DescribeDbInstanceConnectivityResponseBody) SetRequestId(v string) *DescribeDbInstanceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstanceConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
