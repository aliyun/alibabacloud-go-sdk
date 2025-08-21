// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindPurchasedDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUnbindPurchasedDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchUnbindPurchasedDevicesResponseBodyResults) *BatchUnbindPurchasedDevicesResponseBody
	GetResults() []*BatchUnbindPurchasedDevicesResponseBodyResults
}

type BatchUnbindPurchasedDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindPurchasedDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindPurchasedDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUnbindPurchasedDevicesResponseBody) GetResults() []*BatchUnbindPurchasedDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchUnbindPurchasedDevicesResponseBody) SetRequestId(v string) *BatchUnbindPurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBody) SetResults(v []*BatchUnbindPurchasedDevicesResponseBodyResults) *BatchUnbindPurchasedDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchUnbindPurchasedDevicesResponseBodyResults struct {
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// some error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s BatchUnbindPurchasedDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) SetDeviceId(v string) *BatchUnbindPurchasedDevicesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) SetError(v string) *BatchUnbindPurchasedDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
