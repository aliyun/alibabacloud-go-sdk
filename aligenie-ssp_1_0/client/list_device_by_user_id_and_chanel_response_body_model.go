// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdAndChanelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeviceByUserIdAndChanelResponseBody
	GetCode() *int32
	SetMessage(v string) *ListDeviceByUserIdAndChanelResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceByUserIdAndChanelResponseBody
	GetRequestId() *string
	SetResult(v []*ListDeviceByUserIdAndChanelResponseBodyResult) *ListDeviceByUserIdAndChanelResponseBody
	GetResult() []*ListDeviceByUserIdAndChanelResponseBodyResult
}

type ListDeviceByUserIdAndChanelResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// RE***D
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDeviceByUserIdAndChanelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdAndChanelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeviceByUserIdAndChanelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceByUserIdAndChanelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceByUserIdAndChanelResponseBody) GetResult() []*ListDeviceByUserIdAndChanelResponseBodyResult {
	return s.Result
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetCode(v int32) *ListDeviceByUserIdAndChanelResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetMessage(v string) *ListDeviceByUserIdAndChanelResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetRequestId(v string) *ListDeviceByUserIdAndChanelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) SetResult(v []*ListDeviceByUserIdAndChanelResponseBodyResult) *ListDeviceByUserIdAndChanelResponseBody {
	s.Result = v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeviceByUserIdAndChanelResponseBodyResult struct {
	// example:
	//
	// A963*0158
	DeviceOpenId   *string                                                        `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
}

func (s ListDeviceByUserIdAndChanelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) GetDeviceUnionIds() []*ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds {
	return s.DeviceUnionIds
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) SetDeviceOpenId(v string) *ListDeviceByUserIdAndChanelResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) SetDeviceUnionIds(v []*ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) *ListDeviceByUserIdAndChanelResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds struct {
	// example:
	//
	// 1553*B0C3
	DeviceUnionId *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
	// example:
	//
	// 1***2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) GetDeviceUnionId() *string {
	return s.DeviceUnionId
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) SetDeviceUnionId(v string) *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) SetOrganizationId(v string) *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponseBodyResultDeviceUnionIds) Validate() error {
	return dara.Validate(s)
}
