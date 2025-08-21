// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindPurchasedDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchBindPurchasedDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchBindPurchasedDevicesResponseBodyResults) *BatchBindPurchasedDevicesResponseBody
	GetResults() []*BatchBindPurchasedDevicesResponseBodyResults
}

type BatchBindPurchasedDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindPurchasedDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindPurchasedDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchBindPurchasedDevicesResponseBody) GetResults() []*BatchBindPurchasedDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchBindPurchasedDevicesResponseBody) SetRequestId(v string) *BatchBindPurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBody) SetResults(v []*BatchBindPurchasedDevicesResponseBodyResults) *BatchBindPurchasedDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchBindPurchasedDevicesResponseBodyResults struct {
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
	// 238*****380-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s BatchBindPurchasedDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) GetRegion() *string {
	return s.Region
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetDeviceId(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetError(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetGroupId(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.GroupId = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) SetRegion(v string) *BatchBindPurchasedDevicesResponseBodyResults {
	s.Region = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
