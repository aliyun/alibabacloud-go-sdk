// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFraudResultCallBackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *FraudResultCallBackRequest
	GetCertifyId() *string
	SetExtParams(v string) *FraudResultCallBackRequest
	GetExtParams() *string
	SetResultCode(v string) *FraudResultCallBackRequest
	GetResultCode() *string
	SetVerifyDeployEnv(v string) *FraudResultCallBackRequest
	GetVerifyDeployEnv() *string
}

type FraudResultCallBackRequest struct {
	// example:
	//
	// shs2b27333914876c01de4cb22f5841f
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	ExtParams *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	// example:
	//
	// PASS
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// production
	VerifyDeployEnv *string `json:"VerifyDeployEnv,omitempty" xml:"VerifyDeployEnv,omitempty"`
}

func (s FraudResultCallBackRequest) String() string {
	return dara.Prettify(s)
}

func (s FraudResultCallBackRequest) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *FraudResultCallBackRequest) GetExtParams() *string {
	return s.ExtParams
}

func (s *FraudResultCallBackRequest) GetResultCode() *string {
	return s.ResultCode
}

func (s *FraudResultCallBackRequest) GetVerifyDeployEnv() *string {
	return s.VerifyDeployEnv
}

func (s *FraudResultCallBackRequest) SetCertifyId(v string) *FraudResultCallBackRequest {
	s.CertifyId = &v
	return s
}

func (s *FraudResultCallBackRequest) SetExtParams(v string) *FraudResultCallBackRequest {
	s.ExtParams = &v
	return s
}

func (s *FraudResultCallBackRequest) SetResultCode(v string) *FraudResultCallBackRequest {
	s.ResultCode = &v
	return s
}

func (s *FraudResultCallBackRequest) SetVerifyDeployEnv(v string) *FraudResultCallBackRequest {
	s.VerifyDeployEnv = &v
	return s
}

func (s *FraudResultCallBackRequest) Validate() error {
	return dara.Validate(s)
}
