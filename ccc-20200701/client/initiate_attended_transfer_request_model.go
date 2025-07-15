// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateAttendedTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallPriority(v int32) *InitiateAttendedTransferRequest
	GetCallPriority() *int32
	SetDeviceId(v string) *InitiateAttendedTransferRequest
	GetDeviceId() *string
	SetInstanceId(v string) *InitiateAttendedTransferRequest
	GetInstanceId() *string
	SetJobId(v string) *InitiateAttendedTransferRequest
	GetJobId() *string
	SetQueuingOverflowThreshold(v int64) *InitiateAttendedTransferRequest
	GetQueuingOverflowThreshold() *int64
	SetQueuingTimeoutSeconds(v int64) *InitiateAttendedTransferRequest
	GetQueuingTimeoutSeconds() *int64
	SetRoutingType(v string) *InitiateAttendedTransferRequest
	GetRoutingType() *string
	SetStrategyName(v string) *InitiateAttendedTransferRequest
	GetStrategyName() *string
	SetStrategyParams(v string) *InitiateAttendedTransferRequest
	GetStrategyParams() *string
	SetTags(v string) *InitiateAttendedTransferRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *InitiateAttendedTransferRequest
	GetTimeoutSeconds() *int32
	SetTransferee(v string) *InitiateAttendedTransferRequest
	GetTransferee() *string
	SetTransfereeType(v string) *InitiateAttendedTransferRequest
	GetTransfereeType() *string
	SetTransferor(v string) *InitiateAttendedTransferRequest
	GetTransferor() *string
	SetUserId(v string) *InitiateAttendedTransferRequest
	GetUserId() *string
}

type InitiateAttendedTransferRequest struct {
	CallPriority *int32 `json:"CallPriority,omitempty" xml:"CallPriority,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId                    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	QueuingOverflowThreshold *int64  `json:"QueuingOverflowThreshold,omitempty" xml:"QueuingOverflowThreshold,omitempty"`
	QueuingTimeoutSeconds    *int64  `json:"QueuingTimeoutSeconds,omitempty" xml:"QueuingTimeoutSeconds,omitempty"`
	RoutingType              *string `json:"RoutingType,omitempty" xml:"RoutingType,omitempty"`
	StrategyName             *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	StrategyParams           *string `json:"StrategyParams,omitempty" xml:"StrategyParams,omitempty"`
	Tags                     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent2@ccc-test
	Transferee     *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	TransfereeType *string `json:"TransfereeType,omitempty" xml:"TransfereeType,omitempty"`
	Transferor     *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InitiateAttendedTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiateAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferRequest) GetCallPriority() *int32 {
	return s.CallPriority
}

func (s *InitiateAttendedTransferRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InitiateAttendedTransferRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitiateAttendedTransferRequest) GetJobId() *string {
	return s.JobId
}

func (s *InitiateAttendedTransferRequest) GetQueuingOverflowThreshold() *int64 {
	return s.QueuingOverflowThreshold
}

func (s *InitiateAttendedTransferRequest) GetQueuingTimeoutSeconds() *int64 {
	return s.QueuingTimeoutSeconds
}

func (s *InitiateAttendedTransferRequest) GetRoutingType() *string {
	return s.RoutingType
}

func (s *InitiateAttendedTransferRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *InitiateAttendedTransferRequest) GetStrategyParams() *string {
	return s.StrategyParams
}

func (s *InitiateAttendedTransferRequest) GetTags() *string {
	return s.Tags
}

func (s *InitiateAttendedTransferRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *InitiateAttendedTransferRequest) GetTransferee() *string {
	return s.Transferee
}

func (s *InitiateAttendedTransferRequest) GetTransfereeType() *string {
	return s.TransfereeType
}

func (s *InitiateAttendedTransferRequest) GetTransferor() *string {
	return s.Transferor
}

func (s *InitiateAttendedTransferRequest) GetUserId() *string {
	return s.UserId
}

func (s *InitiateAttendedTransferRequest) SetCallPriority(v int32) *InitiateAttendedTransferRequest {
	s.CallPriority = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetDeviceId(v string) *InitiateAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetInstanceId(v string) *InitiateAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetJobId(v string) *InitiateAttendedTransferRequest {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetQueuingOverflowThreshold(v int64) *InitiateAttendedTransferRequest {
	s.QueuingOverflowThreshold = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetQueuingTimeoutSeconds(v int64) *InitiateAttendedTransferRequest {
	s.QueuingTimeoutSeconds = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetRoutingType(v string) *InitiateAttendedTransferRequest {
	s.RoutingType = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetStrategyName(v string) *InitiateAttendedTransferRequest {
	s.StrategyName = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetStrategyParams(v string) *InitiateAttendedTransferRequest {
	s.StrategyParams = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTags(v string) *InitiateAttendedTransferRequest {
	s.Tags = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTimeoutSeconds(v int32) *InitiateAttendedTransferRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferee(v string) *InitiateAttendedTransferRequest {
	s.Transferee = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransfereeType(v string) *InitiateAttendedTransferRequest {
	s.TransfereeType = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferor(v string) *InitiateAttendedTransferRequest {
	s.Transferor = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetUserId(v string) *InitiateAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) Validate() error {
	return dara.Validate(s)
}
