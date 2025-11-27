// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindParentPlatformDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUnbindParentPlatformDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchUnbindParentPlatformDevicesResponseBodyResults) *BatchUnbindParentPlatformDevicesResponseBody
	GetResults() []*BatchUnbindParentPlatformDevicesResponseBodyResults
}

type BatchUnbindParentPlatformDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindParentPlatformDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindParentPlatformDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) GetResults() []*BatchUnbindParentPlatformDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) SetRequestId(v string) *BatchUnbindParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) SetResults(v []*BatchUnbindParentPlatformDevicesResponseBodyResults) *BatchUnbindParentPlatformDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUnbindParentPlatformDevicesResponseBodyResults struct {
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// some error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 361*****212-cn-qingdao
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s BatchUnbindParentPlatformDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetDeviceId(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetError(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) SetParentPlatformId(v string) *BatchUnbindParentPlatformDevicesResponseBodyResults {
	s.ParentPlatformId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
