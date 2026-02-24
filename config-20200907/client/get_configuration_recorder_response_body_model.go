// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigurationRecorderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationRecorder(v *GetConfigurationRecorderResponseBodyConfigurationRecorder) *GetConfigurationRecorderResponseBody
	GetConfigurationRecorder() *GetConfigurationRecorderResponseBodyConfigurationRecorder
	SetRequestId(v string) *GetConfigurationRecorderResponseBody
	GetRequestId() *string
}

type GetConfigurationRecorderResponseBody struct {
	// The resource monitoring information.
	ConfigurationRecorder *GetConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AE43696A-B3AF-5E55-9845-11393127E6D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigurationRecorderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigurationRecorderResponseBody) GetConfigurationRecorder() *GetConfigurationRecorderResponseBodyConfigurationRecorder {
	return s.ConfigurationRecorder
}

func (s *GetConfigurationRecorderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigurationRecorderResponseBody) SetConfigurationRecorder(v *GetConfigurationRecorderResponseBodyConfigurationRecorder) *GetConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *GetConfigurationRecorderResponseBody) SetRequestId(v string) *GetConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigurationRecorderResponseBody) Validate() error {
	if s.ConfigurationRecorder != nil {
		if err := s.ConfigurationRecorder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigurationRecorderResponseBodyConfigurationRecorder struct {
	// The status of resource monitoring. Valid values:
	//
	// - REGISTRABLE: Not registered.
	//
	// - BUILDING: Building.
	//
	// - REGISTERED: Registered.
	//
	// - REBUILDING: Rebuilding.
	//
	// example:
	//
	// REGISTERED
	ConfigurationRecorderStatus *string `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	// A list of monitored resource types.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return dara.Prettify(s)
}

func (s GetConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *GetConfigurationRecorderResponseBodyConfigurationRecorder) GetConfigurationRecorderStatus() *string {
	return s.ConfigurationRecorderStatus
}

func (s *GetConfigurationRecorderResponseBodyConfigurationRecorder) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *GetConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *GetConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *GetConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

func (s *GetConfigurationRecorderResponseBodyConfigurationRecorder) Validate() error {
	return dara.Validate(s)
}
