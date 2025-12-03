// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVMDeployOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployOrder(v *GetVMDeployOrderResponseBodyDeployOrder) *GetVMDeployOrderResponseBody
	GetDeployOrder() *GetVMDeployOrderResponseBodyDeployOrder
	SetErrorCode(v string) *GetVMDeployOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetVMDeployOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetVMDeployOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVMDeployOrderResponseBody
	GetSuccess() *bool
}

type GetVMDeployOrderResponseBody struct {
	DeployOrder *GetVMDeployOrderResponseBodyDeployOrder `json:"deployOrder,omitempty" xml:"deployOrder,omitempty" type:"Struct"`
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVMDeployOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBody) GetDeployOrder() *GetVMDeployOrderResponseBodyDeployOrder {
	return s.DeployOrder
}

func (s *GetVMDeployOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVMDeployOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVMDeployOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVMDeployOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVMDeployOrderResponseBody) SetDeployOrder(v *GetVMDeployOrderResponseBodyDeployOrder) *GetVMDeployOrderResponseBody {
	s.DeployOrder = v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetErrorCode(v string) *GetVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetErrorMessage(v string) *GetVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetRequestId(v string) *GetVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetSuccess(v bool) *GetVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) Validate() error {
	if s.DeployOrder != nil {
		if err := s.DeployOrder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVMDeployOrderResponseBodyDeployOrder struct {
	Actions []*GetVMDeployOrderResponseBodyDeployOrderActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// 111111111111
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// ssaassa
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 2
	CurrentBatch      *int32                                                    `json:"currentBatch,omitempty" xml:"currentBatch,omitempty"`
	DeployMachineInfo *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo `json:"deployMachineInfo,omitempty" xml:"deployMachineInfo,omitempty" type:"Struct"`
	// example:
	//
	// 11111
	DeployOrderId *string `json:"deployOrderId,omitempty" xml:"deployOrderId,omitempty"`
	ExceptionCode *string `json:"exceptionCode,omitempty" xml:"exceptionCode,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 3
	TotalBatch *int32 `json:"totalBatch,omitempty" xml:"totalBatch,omitempty"`
	// example:
	//
	// 11111111111
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrder) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrder) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetActions() []*GetVMDeployOrderResponseBodyDeployOrderActions {
	return s.Actions
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetCreator() *string {
	return s.Creator
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetCurrentBatch() *int32 {
	return s.CurrentBatch
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetDeployMachineInfo() *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	return s.DeployMachineInfo
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetDeployOrderId() *string {
	return s.DeployOrderId
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetExceptionCode() *string {
	return s.ExceptionCode
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetStatus() *string {
	return s.Status
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetTotalBatch() *int32 {
	return s.TotalBatch
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetActions(v []*GetVMDeployOrderResponseBodyDeployOrderActions) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Actions = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCreateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrder {
	s.CreateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCreator(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Creator = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCurrentBatch(v int32) *GetVMDeployOrderResponseBodyDeployOrder {
	s.CurrentBatch = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetDeployMachineInfo(v *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) *GetVMDeployOrderResponseBodyDeployOrder {
	s.DeployMachineInfo = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetDeployOrderId(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.DeployOrderId = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetExceptionCode(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.ExceptionCode = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetStatus(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Status = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetTotalBatch(v int32) *GetVMDeployOrderResponseBodyDeployOrder {
	s.TotalBatch = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetUpdateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrder {
	s.UpdateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeployMachineInfo != nil {
		if err := s.DeployMachineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVMDeployOrderResponseBodyDeployOrderActions struct {
	// example:
	//
	// true
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// {}
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// StopVMDeployOrder
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderActions) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderActions) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) GetDisable() *bool {
	return s.Disable
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) GetParams() interface{} {
	return s.Params
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) GetType() *string {
	return s.Type
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetDisable(v bool) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Disable = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetParams(v interface{}) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Type = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) Validate() error {
	return dara.Validate(s)
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo struct {
	// example:
	//
	// 11
	BatchNum       *int32                                                                    `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	DeployMachines []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines `json:"deployMachines,omitempty" xml:"deployMachines,omitempty" type:"Repeated"`
	// example:
	//
	// 1111
	HostGroupId *int64 `json:"hostGroupId,omitempty" xml:"hostGroupId,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) GetBatchNum() *int32 {
	return s.BatchNum
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) GetDeployMachines() []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	return s.DeployMachines
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) GetHostGroupId() *int64 {
	return s.HostGroupId
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetBatchNum(v int32) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.BatchNum = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetDeployMachines(v []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.DeployMachines = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetHostGroupId(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.HostGroupId = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) Validate() error {
	if s.DeployMachines != nil {
		for _, item := range s.DeployMachines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines struct {
	Actions []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	BatchNum *int32 `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	// example:
	//
	// ok
	ClientStatus *string `json:"clientStatus,omitempty" xml:"clientStatus,omitempty"`
	// example:
	//
	// 1111111111
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// example:
	//
	// sasssasa
	MachineSn *string `json:"machineSn,omitempty" xml:"machineSn,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 11111111
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetActions() []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	return s.Actions
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetBatchNum() *int32 {
	return s.BatchNum
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetIp() *string {
	return s.Ip
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetMachineSn() *string {
	return s.MachineSn
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetStatus() *string {
	return s.Status
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetActions(v []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Actions = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetBatchNum(v int32) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.BatchNum = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetClientStatus(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.ClientStatus = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetCreateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.CreateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetIp(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Ip = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetMachineSn(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.MachineSn = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetStatus(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Status = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetUpdateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.UpdateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions struct {
	// example:
	//
	// true
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// {}
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// RetryVMDeployMachine
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) GetDisable() *bool {
	return s.Disable
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) GetParams() interface{} {
	return s.Params
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) GetType() *string {
	return s.Type
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetDisable(v bool) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Disable = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetParams(v interface{}) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Type = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) Validate() error {
	return dara.Validate(s)
}
