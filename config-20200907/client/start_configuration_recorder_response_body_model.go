// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigurationRecorderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationRecorder(v *StartConfigurationRecorderResponseBodyConfigurationRecorder) *StartConfigurationRecorderResponseBody
	GetConfigurationRecorder() *StartConfigurationRecorderResponseBodyConfigurationRecorder
	SetRequestId(v string) *StartConfigurationRecorderResponseBody
	GetRequestId() *string
}

type StartConfigurationRecorderResponseBody struct {
	ConfigurationRecorder *StartConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	RequestId             *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartConfigurationRecorderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponseBody) GetConfigurationRecorder() *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	return s.ConfigurationRecorder
}

func (s *StartConfigurationRecorderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartConfigurationRecorderResponseBody) SetConfigurationRecorder(v *StartConfigurationRecorderResponseBodyConfigurationRecorder) *StartConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *StartConfigurationRecorderResponseBody) SetRequestId(v string) *StartConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartConfigurationRecorderResponseBody) Validate() error {
	if s.ConfigurationRecorder != nil {
		if err := s.ConfigurationRecorder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartConfigurationRecorderResponseBodyConfigurationRecorder struct {
	ConfigurationRecorderStatus *string   `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	ResourceTypes               []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s StartConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return dara.Prettify(s)
}

func (s StartConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) GetConfigurationRecorderStatus() *string {
	return s.ConfigurationRecorderStatus
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) Validate() error {
	return dara.Validate(s)
}
