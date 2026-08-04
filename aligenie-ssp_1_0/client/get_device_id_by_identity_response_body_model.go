// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceIdByIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceIdByIdentityResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceIdByIdentityResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceIdByIdentityResponseBody
	GetRequestId() *string
	SetResult(v *GetDeviceIdByIdentityResponseBodyResult) *GetDeviceIdByIdentityResponseBody
	GetResult() *GetDeviceIdByIdentityResponseBodyResult
}

type GetDeviceIdByIdentityResponseBody struct {
	// The error code returned. A value of 200 indicates that the invocation succeeded.
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
	// Request ID.
	//
	// example:
	//
	// 0EC7DA****A0726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed information returned.
	Result *GetDeviceIdByIdentityResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceIdByIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceIdByIdentityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceIdByIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceIdByIdentityResponseBody) GetResult() *GetDeviceIdByIdentityResponseBodyResult {
	return s.Result
}

func (s *GetDeviceIdByIdentityResponseBody) SetCode(v int32) *GetDeviceIdByIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetMessage(v string) *GetDeviceIdByIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetRequestId(v string) *GetDeviceIdByIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) SetResult(v *GetDeviceIdByIdentityResponseBodyResult) *GetDeviceIdByIdentityResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceIdByIdentityResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceIdByIdentityResponseBodyResult struct {
	// The openId corresponding to the device.
	//
	// example:
	//
	// A963*0158
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// Organization ID and UnionId information corresponding to the device.
	DeviceUnionIds []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s GetDeviceIdByIdentityResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *GetDeviceIdByIdentityResponseBodyResult) GetDeviceUnionIds() []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	return s.DeviceUnionIds
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceOpenId(v string) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResult) SetDeviceUnionIds(v []*GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) *GetDeviceIdByIdentityResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResult) Validate() error {
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

type GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds struct {
	// The UnionId of the device.
	//
	// example:
	//
	// 1553*B0C3
	DeviceUnionId *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	// Organization ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) GetDeviceUnionId() *string {
	return s.DeviceUnionId
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceIdByIdentityResponseBodyResultDeviceUnionIds) Validate() error {
	return dara.Validate(s)
}
