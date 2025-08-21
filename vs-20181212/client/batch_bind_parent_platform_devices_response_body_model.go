// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindParentPlatformDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchBindParentPlatformDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchBindParentPlatformDevicesResponseBodyResults) *BatchBindParentPlatformDevicesResponseBody
	GetResults() []*BatchBindParentPlatformDevicesResponseBodyResults
}

type BatchBindParentPlatformDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindParentPlatformDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindParentPlatformDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchBindParentPlatformDevicesResponseBody) GetResults() []*BatchBindParentPlatformDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchBindParentPlatformDevicesResponseBody) SetRequestId(v string) *BatchBindParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBody) SetResults(v []*BatchBindParentPlatformDevicesResponseBodyResults) *BatchBindParentPlatformDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchBindParentPlatformDevicesResponseBodyResults struct {
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

func (s BatchBindParentPlatformDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetDeviceId(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetError(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) SetParentPlatformId(v string) *BatchBindParentPlatformDevicesResponseBodyResults {
	s.ParentPlatformId = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
