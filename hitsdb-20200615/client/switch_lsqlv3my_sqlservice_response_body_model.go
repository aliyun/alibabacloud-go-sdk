// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchLSQLV3MySQLServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SwitchLSQLV3MySQLServiceResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *SwitchLSQLV3MySQLServiceResponseBody
	GetRequestId() *string
}

type SwitchLSQLV3MySQLServiceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) SetAccessDeniedDetail(v string) *SwitchLSQLV3MySQLServiceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) SetRequestId(v string) *SwitchLSQLV3MySQLServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
