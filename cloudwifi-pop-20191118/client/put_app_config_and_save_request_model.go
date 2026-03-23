// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAppConfigAndSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *PutAppConfigAndSaveRequest
	GetApMac() *string
	SetAppCode(v string) *PutAppConfigAndSaveRequest
	GetAppCode() *string
	SetAppName(v string) *PutAppConfigAndSaveRequest
	GetAppName() *string
	SetConfigureType(v string) *PutAppConfigAndSaveRequest
	GetConfigureType() *string
	SetCurrentTime(v int64) *PutAppConfigAndSaveRequest
	GetCurrentTime() *int64
	SetData(v string) *PutAppConfigAndSaveRequest
	GetData() *string
}

type PutAppConfigAndSaveRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	ConfigureType *string `json:"ConfigureType,omitempty" xml:"ConfigureType,omitempty"`
	// This parameter is required.
	CurrentTime *int64 `json:"CurrentTime,omitempty" xml:"CurrentTime,omitempty"`
	// This parameter is required.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s PutAppConfigAndSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAppConfigAndSaveRequest) GoString() string {
	return s.String()
}

func (s *PutAppConfigAndSaveRequest) GetApMac() *string {
	return s.ApMac
}

func (s *PutAppConfigAndSaveRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *PutAppConfigAndSaveRequest) GetAppName() *string {
	return s.AppName
}

func (s *PutAppConfigAndSaveRequest) GetConfigureType() *string {
	return s.ConfigureType
}

func (s *PutAppConfigAndSaveRequest) GetCurrentTime() *int64 {
	return s.CurrentTime
}

func (s *PutAppConfigAndSaveRequest) GetData() *string {
	return s.Data
}

func (s *PutAppConfigAndSaveRequest) SetApMac(v string) *PutAppConfigAndSaveRequest {
	s.ApMac = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) SetAppCode(v string) *PutAppConfigAndSaveRequest {
	s.AppCode = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) SetAppName(v string) *PutAppConfigAndSaveRequest {
	s.AppName = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) SetConfigureType(v string) *PutAppConfigAndSaveRequest {
	s.ConfigureType = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) SetCurrentTime(v int64) *PutAppConfigAndSaveRequest {
	s.CurrentTime = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) SetData(v string) *PutAppConfigAndSaveRequest {
	s.Data = &v
	return s
}

func (s *PutAppConfigAndSaveRequest) Validate() error {
	return dara.Validate(s)
}
