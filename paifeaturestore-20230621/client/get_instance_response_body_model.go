// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureDBInfo(v *GetInstanceResponseBodyFeatureDBInfo) *GetInstanceResponseBody
	GetFeatureDBInfo() *GetInstanceResponseBodyFeatureDBInfo
	SetFeatureDBInstanceInfo(v *GetInstanceResponseBodyFeatureDBInstanceInfo) *GetInstanceResponseBody
	GetFeatureDBInstanceInfo() *GetInstanceResponseBodyFeatureDBInstanceInfo
	SetGmtCreateTime(v string) *GetInstanceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceResponseBody
	GetGmtModifiedTime() *string
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetProgress(v float64) *GetInstanceResponseBody
	GetProgress() *float64
	SetRegionId(v string) *GetInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetInstanceResponseBody
	GetStatus() *string
	SetType(v string) *GetInstanceResponseBody
	GetType() *string
}

type GetInstanceResponseBody struct {
	// The information about the FeatureDB instance.
	FeatureDBInfo *GetInstanceResponseBodyFeatureDBInfo `json:"FeatureDBInfo,omitempty" xml:"FeatureDBInfo,omitempty" type:"Struct"`
	// Deprecated
	//
	// The information about the FeatureDB instance.
	FeatureDBInstanceInfo *GetInstanceResponseBodyFeatureDBInstanceInfo `json:"FeatureDBInstanceInfo,omitempty" xml:"FeatureDBInstanceInfo,omitempty" type:"Struct"`
	// The time when the instance was created. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the instance was last updated. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The message that is returned for the instance provisioning.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress of instance provisioning.
	//
	// example:
	//
	// 0.8
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C5B1511-8A5B-59C3-90AF-513F9210E882
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the instance. Valid values:
	//
	// - Initializing: The instance is being initialized.
	//
	// - Running: The instance is running.
	//
	// - Failure: The instance fails to be created.
	//
	// - Stopped: The instance is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetFeatureDBInfo() *GetInstanceResponseBodyFeatureDBInfo {
	return s.FeatureDBInfo
}

func (s *GetInstanceResponseBody) GetFeatureDBInstanceInfo() *GetInstanceResponseBodyFeatureDBInstanceInfo {
	return s.FeatureDBInstanceInfo
}

func (s *GetInstanceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetProgress() *float64 {
	return s.Progress
}

func (s *GetInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBody) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBody) SetFeatureDBInfo(v *GetInstanceResponseBodyFeatureDBInfo) *GetInstanceResponseBody {
	s.FeatureDBInfo = v
	return s
}

func (s *GetInstanceResponseBody) SetFeatureDBInstanceInfo(v *GetInstanceResponseBodyFeatureDBInstanceInfo) *GetInstanceResponseBody {
	s.FeatureDBInstanceInfo = v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetProgress(v float64) *GetInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetType(v string) *GetInstanceResponseBody {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.FeatureDBInfo != nil {
		if err := s.FeatureDBInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FeatureDBInstanceInfo != nil {
		if err := s.FeatureDBInstanceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyFeatureDBInfo struct {
	// The status of the instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyFeatureDBInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyFeatureDBInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyFeatureDBInfo) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyFeatureDBInfo) SetStatus(v string) *GetInstanceResponseBodyFeatureDBInfo {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyFeatureDBInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyFeatureDBInstanceInfo struct {
	// The status of the instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyFeatureDBInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyFeatureDBInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyFeatureDBInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyFeatureDBInstanceInfo) SetStatus(v string) *GetInstanceResponseBodyFeatureDBInstanceInfo {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyFeatureDBInstanceInfo) Validate() error {
	return dara.Validate(s)
}
