// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryPersistenceDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPersistenceDay(v int32) *UpdateJenkinsImageRegistryPersistenceDayRequest
	GetPersistenceDay() *int32
	SetRegistryId(v int64) *UpdateJenkinsImageRegistryPersistenceDayRequest
	GetRegistryId() *int64
	SetSourceIp(v string) *UpdateJenkinsImageRegistryPersistenceDayRequest
	GetSourceIp() *string
}

type UpdateJenkinsImageRegistryPersistenceDayRequest struct {
	// The retention period. Unit: days.
	//
	// example:
	//
	// 30
	PersistenceDay *int32 `json:"PersistenceDay,omitempty" xml:"PersistenceDay,omitempty"`
	// The ID of the image repository.
	//
	// > You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to query the IDs of image repositories.
	//
	// example:
	//
	// 25363
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 106.11.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateJenkinsImageRegistryPersistenceDayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryPersistenceDayRequest) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) GetPersistenceDay() *int32 {
	return s.PersistenceDay
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) SetPersistenceDay(v int32) *UpdateJenkinsImageRegistryPersistenceDayRequest {
	s.PersistenceDay = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) SetRegistryId(v int64) *UpdateJenkinsImageRegistryPersistenceDayRequest {
	s.RegistryId = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) SetSourceIp(v string) *UpdateJenkinsImageRegistryPersistenceDayRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateJenkinsImageRegistryPersistenceDayRequest) Validate() error {
	return dara.Validate(s)
}
