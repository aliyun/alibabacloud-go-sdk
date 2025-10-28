// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJvmConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJvmConfigurationResponseBody
	GetCode() *int32
	SetJvmConfiguration(v *GetJvmConfigurationResponseBodyJvmConfiguration) *GetJvmConfigurationResponseBody
	GetJvmConfiguration() *GetJvmConfigurationResponseBodyJvmConfiguration
	SetMessage(v string) *GetJvmConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJvmConfigurationResponseBody
	GetRequestId() *string
}

type GetJvmConfigurationResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The JVM configuration of the application or instance group.
	JvmConfiguration *GetJvmConfigurationResponseBodyJvmConfiguration `json:"JvmConfiguration,omitempty" xml:"JvmConfiguration,omitempty" type:"Struct"`
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
	// 3F43-F34V-0VCD***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJvmConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJvmConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJvmConfigurationResponseBody) GetJvmConfiguration() *GetJvmConfigurationResponseBodyJvmConfiguration {
	return s.JvmConfiguration
}

func (s *GetJvmConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJvmConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJvmConfigurationResponseBody) SetCode(v int32) *GetJvmConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetJvmConfiguration(v *GetJvmConfigurationResponseBodyJvmConfiguration) *GetJvmConfigurationResponseBody {
	s.JvmConfiguration = v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetMessage(v string) *GetJvmConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetRequestId(v string) *GetJvmConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) Validate() error {
	if s.JvmConfiguration != nil {
		if err := s.JvmConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJvmConfigurationResponseBodyJvmConfiguration struct {
	// The maximum size of the heap memory. Unit: MB.
	//
	// example:
	//
	// 1000
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
	// The custom parameter.
	//
	// example:
	//
	// -XX:+UseConcMarkSweepGC -XX:-UseParNewGC
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s GetJvmConfigurationResponseBodyJvmConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetJvmConfigurationResponseBodyJvmConfiguration) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) GetMaxHeapSize() *int32 {
	return s.MaxHeapSize
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) GetMaxPermSize() *int32 {
	return s.MaxPermSize
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) GetMinHeapSize() *int32 {
	return s.MinHeapSize
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) GetOptions() *string {
	return s.Options
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMaxHeapSize(v int32) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxHeapSize = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMaxPermSize(v int32) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxPermSize = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMinHeapSize(v int32) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MinHeapSize = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetOptions(v string) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.Options = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) Validate() error {
	return dara.Validate(s)
}
