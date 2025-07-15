// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnoseId(v string) *DiagnoseVpnGatewayResponseBody
	GetDiagnoseId() *string
	SetRequestId(v string) *DiagnoseVpnGatewayResponseBody
	GetRequestId() *string
}

type DiagnoseVpnGatewayResponseBody struct {
	// The diagnostic ID.
	//
	// After a diagnostic ID is returned, you can call [GetVpnGatewayDiagnoseResult](https://help.aliyun.com/document_detail/2521963.html) to query the diagnostic report.
	//
	// example:
	//
	// vpndgn-uf6kuxbe3iv028k3s****
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DiagnoseVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnGatewayResponseBody) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DiagnoseVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiagnoseVpnGatewayResponseBody) SetDiagnoseId(v string) *DiagnoseVpnGatewayResponseBody {
	s.DiagnoseId = &v
	return s
}

func (s *DiagnoseVpnGatewayResponseBody) SetRequestId(v string) *DiagnoseVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
