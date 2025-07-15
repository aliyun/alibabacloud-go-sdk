// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowPackageData(v string) *ImportContactFlowRequest
	GetFlowPackageData() *string
	SetInstanceId(v string) *ImportContactFlowRequest
	GetInstanceId() *string
	SetRequestId(v string) *ImportContactFlowRequest
	GetRequestId() *string
}

type ImportContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	FlowPackageData *string `json:"FlowPackageData,omitempty" xml:"FlowPackageData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportContactFlowRequest) GoString() string {
	return s.String()
}

func (s *ImportContactFlowRequest) GetFlowPackageData() *string {
	return s.FlowPackageData
}

func (s *ImportContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportContactFlowRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportContactFlowRequest) SetFlowPackageData(v string) *ImportContactFlowRequest {
	s.FlowPackageData = &v
	return s
}

func (s *ImportContactFlowRequest) SetInstanceId(v string) *ImportContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportContactFlowRequest) SetRequestId(v string) *ImportContactFlowRequest {
	s.RequestId = &v
	return s
}

func (s *ImportContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
