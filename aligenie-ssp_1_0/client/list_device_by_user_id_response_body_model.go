// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeviceByUserIdResponseBody
	GetCode() *int32
	SetMessage(v string) *ListDeviceByUserIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceByUserIdResponseBody
	GetRequestId() *string
	SetResult(v []*ListDeviceByUserIdResponseBodyResult) *ListDeviceByUserIdResponseBody
	GetResult() []*ListDeviceByUserIdResponseBodyResult
}

type ListDeviceByUserIdResponseBody struct {
	// The returned error code, where 200 indicates that the invocation succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return result of invoking this API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed information returned.
	Result []*ListDeviceByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeviceByUserIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceByUserIdResponseBody) GetResult() []*ListDeviceByUserIdResponseBodyResult {
	return s.Result
}

func (s *ListDeviceByUserIdResponseBody) SetCode(v int32) *ListDeviceByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetMessage(v string) *ListDeviceByUserIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetRequestId(v string) *ListDeviceByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBody) SetResult(v []*ListDeviceByUserIdResponseBodyResult) *ListDeviceByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *ListDeviceByUserIdResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceByUserIdResponseBodyResult struct {
	// The openId corresponding to the Device Information.
	//
	// example:
	//
	// A963*0158
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// Organization ID and UnionId information corresponding to the device.
	DeviceUnionIds []*ListDeviceByUserIdResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *ListDeviceByUserIdResponseBodyResult) GetDeviceUnionIds() []*ListDeviceByUserIdResponseBodyResultDeviceUnionIds {
	return s.DeviceUnionIds
}

func (s *ListDeviceByUserIdResponseBodyResult) SetDeviceOpenId(v string) *ListDeviceByUserIdResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResult) SetDeviceUnionIds(v []*ListDeviceByUserIdResponseBodyResultDeviceUnionIds) *ListDeviceByUserIdResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResult) Validate() error {
	if s.DeviceUnionIds != nil {
		for _, item := range s.DeviceUnionIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceByUserIdResponseBodyResultDeviceUnionIds struct {
	// The UnionId of the device.
	//
	// example:
	//
	// 1553*B0C3
	DeviceUnionId *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdResponseBodyResultDeviceUnionIds) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) GetDeviceUnionId() *string {
	return s.DeviceUnionId
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *ListDeviceByUserIdResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *ListDeviceByUserIdResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *ListDeviceByUserIdResponseBodyResultDeviceUnionIds) Validate() error {
	return dara.Validate(s)
}
