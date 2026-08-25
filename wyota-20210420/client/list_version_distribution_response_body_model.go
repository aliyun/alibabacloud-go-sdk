// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVersionDistributionResponseBody
	GetCode() *string
	SetData(v []*ListVersionDistributionResponseBodyData) *ListVersionDistributionResponseBody
	GetData() []*ListVersionDistributionResponseBodyData
	SetHttpStatusCode(v int32) *ListVersionDistributionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVersionDistributionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVersionDistributionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVersionDistributionResponseBody
	GetSuccess() *bool
}

type ListVersionDistributionResponseBody struct {
	// The status code. A value of 200 is returned if the call is successful. An error code is returned if the call fails.
	//
	// example:
	//
	// PARAMETER_MISSING
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of version distribution information.
	Data []*ListVersionDistributionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. This parameter is empty if the call is successful.
	//
	// example:
	//
	// parameter missing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVersionDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVersionDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionDistributionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVersionDistributionResponseBody) GetData() []*ListVersionDistributionResponseBodyData {
	return s.Data
}

func (s *ListVersionDistributionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVersionDistributionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVersionDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVersionDistributionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVersionDistributionResponseBody) SetCode(v string) *ListVersionDistributionResponseBody {
	s.Code = &v
	return s
}

func (s *ListVersionDistributionResponseBody) SetData(v []*ListVersionDistributionResponseBodyData) *ListVersionDistributionResponseBody {
	s.Data = v
	return s
}

func (s *ListVersionDistributionResponseBody) SetHttpStatusCode(v int32) *ListVersionDistributionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVersionDistributionResponseBody) SetMessage(v string) *ListVersionDistributionResponseBody {
	s.Message = &v
	return s
}

func (s *ListVersionDistributionResponseBody) SetRequestId(v string) *ListVersionDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionDistributionResponseBody) SetSuccess(v bool) *ListVersionDistributionResponseBody {
	s.Success = &v
	return s
}

func (s *ListVersionDistributionResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVersionDistributionResponseBodyData struct {
	// The number of terminals corresponding to this version.
	//
	// example:
	//
	// 60
	DeviceCount *int64 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// The version percentage. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.6
	Percentage *float64 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The version number.
	//
	// example:
	//
	// 2.3.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListVersionDistributionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVersionDistributionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVersionDistributionResponseBodyData) GetDeviceCount() *int64 {
	return s.DeviceCount
}

func (s *ListVersionDistributionResponseBodyData) GetPercentage() *float64 {
	return s.Percentage
}

func (s *ListVersionDistributionResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListVersionDistributionResponseBodyData) SetDeviceCount(v int64) *ListVersionDistributionResponseBodyData {
	s.DeviceCount = &v
	return s
}

func (s *ListVersionDistributionResponseBodyData) SetPercentage(v float64) *ListVersionDistributionResponseBodyData {
	s.Percentage = &v
	return s
}

func (s *ListVersionDistributionResponseBodyData) SetVersion(v string) *ListVersionDistributionResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListVersionDistributionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
