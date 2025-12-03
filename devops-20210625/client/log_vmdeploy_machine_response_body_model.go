// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogVMDeployMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployMachineLog(v *LogVMDeployMachineResponseBodyDeployMachineLog) *LogVMDeployMachineResponseBody
	GetDeployMachineLog() *LogVMDeployMachineResponseBodyDeployMachineLog
	SetErrorCode(v string) *LogVMDeployMachineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *LogVMDeployMachineResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *LogVMDeployMachineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LogVMDeployMachineResponseBody
	GetSuccess() *bool
}

type LogVMDeployMachineResponseBody struct {
	DeployMachineLog *LogVMDeployMachineResponseBodyDeployMachineLog `json:"deployMachineLog,omitempty" xml:"deployMachineLog,omitempty" type:"Struct"`
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

func (s LogVMDeployMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LogVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponseBody) GetDeployMachineLog() *LogVMDeployMachineResponseBodyDeployMachineLog {
	return s.DeployMachineLog
}

func (s *LogVMDeployMachineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LogVMDeployMachineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *LogVMDeployMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LogVMDeployMachineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LogVMDeployMachineResponseBody) SetDeployMachineLog(v *LogVMDeployMachineResponseBodyDeployMachineLog) *LogVMDeployMachineResponseBody {
	s.DeployMachineLog = v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetErrorCode(v string) *LogVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetErrorMessage(v string) *LogVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetRequestId(v string) *LogVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetSuccess(v bool) *LogVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) Validate() error {
	if s.DeployMachineLog != nil {
		if err := s.DeployMachineLog.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LogVMDeployMachineResponseBodyDeployMachineLog struct {
	// example:
	//
	// cn-hangzhou
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	// example:
	//
	// 11111111111
	DeployBeginTime *int64 `json:"deployBeginTime,omitempty" xml:"deployBeginTime,omitempty"`
	// example:
	//
	// 12222222
	DeployEndTime *int64 `json:"deployEndTime,omitempty" xml:"deployEndTime,omitempty"`
	// example:
	//
	// success
	DeployLog *string `json:"deployLog,omitempty" xml:"deployLog,omitempty"`
	// example:
	//
	// /tmp/log
	DeployLogPath *string `json:"deployLogPath,omitempty" xml:"deployLogPath,omitempty"`
}

func (s LogVMDeployMachineResponseBodyDeployMachineLog) String() string {
	return dara.Prettify(s)
}

func (s LogVMDeployMachineResponseBodyDeployMachineLog) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) GetAliyunRegion() *string {
	return s.AliyunRegion
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) GetDeployBeginTime() *int64 {
	return s.DeployBeginTime
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) GetDeployEndTime() *int64 {
	return s.DeployEndTime
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) GetDeployLog() *string {
	return s.DeployLog
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) GetDeployLogPath() *string {
	return s.DeployLogPath
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetAliyunRegion(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.AliyunRegion = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployBeginTime(v int64) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployBeginTime = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployEndTime(v int64) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployEndTime = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployLog(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployLog = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployLogPath(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployLogPath = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) Validate() error {
	return dara.Validate(s)
}
