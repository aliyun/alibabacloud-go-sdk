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
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of results.
	Results []*BatchBindPurchasedDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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

type BatchBindPurchasedDevicesResponseBodyResults struct {
	// The ID of the device.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The error message for the device. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// some error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// Space ID.
	//
	// example:
	//
	// 238*****380-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The region.
	//
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
