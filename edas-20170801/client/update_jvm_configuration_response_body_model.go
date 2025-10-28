// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJvmConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateJvmConfigurationResponseBody
	GetCode() *int32
	SetJvmConfiguration(v *UpdateJvmConfigurationResponseBodyJvmConfiguration) *UpdateJvmConfigurationResponseBody
	GetJvmConfiguration() *UpdateJvmConfigurationResponseBodyJvmConfiguration
	SetMessage(v string) *UpdateJvmConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJvmConfigurationResponseBody
	GetRequestId() *string
}

type UpdateJvmConfigurationResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The settings of the JVM parameters.
	JvmConfiguration *UpdateJvmConfigurationResponseBodyJvmConfiguration `json:"JvmConfiguration,omitempty" xml:"JvmConfiguration,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-********************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJvmConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJvmConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateJvmConfigurationResponseBody) GetJvmConfiguration() *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	return s.JvmConfiguration
}

func (s *UpdateJvmConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJvmConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJvmConfigurationResponseBody) SetCode(v int32) *UpdateJvmConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetJvmConfiguration(v *UpdateJvmConfigurationResponseBodyJvmConfiguration) *UpdateJvmConfigurationResponseBody {
	s.JvmConfiguration = v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetMessage(v string) *UpdateJvmConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetRequestId(v string) *UpdateJvmConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) Validate() error {
	if s.JvmConfiguration != nil {
		if err := s.JvmConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateJvmConfigurationResponseBodyJvmConfiguration struct {
	// The maximum size of the heap memory. Unit: MB.
	//
	// example:
	//
	// 500
	MaxHeapSize *int32 `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty"`
	// The size of the permanent generation heap memory. Unit: MB.
	//
	// example:
	//
	// 1000
	MaxPermSize *int32 `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty"`
	// The initial size of the heap memory. Unit: MB.
	//
	// example:
	//
	// 500
	MinHeapSize *int32 `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty"`
	// The optional parameters.
	//
	// example:
	//
	// ”“
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateJvmConfigurationResponseBodyJvmConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateJvmConfigurationResponseBodyJvmConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) GetMaxHeapSize() *int32 {
	return s.MaxHeapSize
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) GetMaxPermSize() *int32 {
	return s.MaxPermSize
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) GetMinHeapSize() *int32 {
	return s.MinHeapSize
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) GetOptions() *string {
	return s.Options
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMaxHeapSize(v int32) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMaxPermSize(v int32) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxPermSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMinHeapSize(v int32) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MinHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetOptions(v string) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.Options = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) Validate() error {
	return dara.Validate(s)
}
