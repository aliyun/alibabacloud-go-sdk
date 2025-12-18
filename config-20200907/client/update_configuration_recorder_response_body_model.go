// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigurationRecorderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationRecorder(v *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) *UpdateConfigurationRecorderResponseBody
	GetConfigurationRecorder() *UpdateConfigurationRecorderResponseBodyConfigurationRecorder
	SetRequestId(v string) *UpdateConfigurationRecorderResponseBody
	GetRequestId() *string
}

type UpdateConfigurationRecorderResponseBody struct {
	// The details of the configuration recorder.
	ConfigurationRecorder *UpdateConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2A1D5095-D1F9-56B6-A712-EA495608E79A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigurationRecorderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationRecorderResponseBody) GetConfigurationRecorder() *UpdateConfigurationRecorderResponseBodyConfigurationRecorder {
	return s.ConfigurationRecorder
}

func (s *UpdateConfigurationRecorderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigurationRecorderResponseBody) SetConfigurationRecorder(v *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) *UpdateConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *UpdateConfigurationRecorderResponseBody) SetRequestId(v string) *UpdateConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigurationRecorderResponseBody) Validate() error {
	if s.ConfigurationRecorder != nil {
		if err := s.ConfigurationRecorder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConfigurationRecorderResponseBodyConfigurationRecorder struct {
	// The status of the configuration recorder. Valid values:
	//
	// 	- REGISTRABLE: The configuration recorder has not been registered.
	//
	// 	- BUILDING: The configuration recorder is being deployed.
	//
	// 	- REGISTERED: The configuration recorder has been registered.
	//
	// 	- REBUILDING: The configuration recorder is being redeployed.
	//
	// example:
	//
	// REGISTERED
	ConfigurationRecorderStatus *string `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	// The types of the resources that are monitored by Cloud Config.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s UpdateConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) GetConfigurationRecorderStatus() *string {
	return s.ConfigurationRecorderStatus
}

func (s *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *UpdateConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *UpdateConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

func (s *UpdateConfigurationRecorderResponseBodyConfigurationRecorder) Validate() error {
	return dara.Validate(s)
}
