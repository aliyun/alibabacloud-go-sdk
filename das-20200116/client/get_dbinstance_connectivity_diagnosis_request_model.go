// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceConnectivityDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetDBInstanceConnectivityDiagnosisRequest
	GetInstanceId() *string
	SetSrcIp(v string) *GetDBInstanceConnectivityDiagnosisRequest
	GetSrcIp() *string
}

type GetDBInstanceConnectivityDiagnosisRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.110.180.62
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s GetDBInstanceConnectivityDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceConnectivityDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *GetDBInstanceConnectivityDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDBInstanceConnectivityDiagnosisRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *GetDBInstanceConnectivityDiagnosisRequest) SetInstanceId(v string) *GetDBInstanceConnectivityDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisRequest) SetSrcIp(v string) *GetDBInstanceConnectivityDiagnosisRequest {
	s.SrcIp = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
