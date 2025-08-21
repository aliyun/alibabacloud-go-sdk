// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAuthCodeBindForExtResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CheckAuthCodeBindForExtResponseBody
	GetCode() *int32
	SetMessage(v string) *CheckAuthCodeBindForExtResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAuthCodeBindForExtResponseBody
	GetRequestId() *string
	SetResult(v *CheckAuthCodeBindForExtResponseBodyResult) *CheckAuthCodeBindForExtResponseBody
	GetResult() *CheckAuthCodeBindForExtResponseBodyResult
}

type CheckAuthCodeBindForExtResponseBody struct {
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
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckAuthCodeBindForExtResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CheckAuthCodeBindForExtResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAuthCodeBindForExtResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAuthCodeBindForExtResponseBody) GetResult() *CheckAuthCodeBindForExtResponseBodyResult {
	return s.Result
}

func (s *CheckAuthCodeBindForExtResponseBody) SetCode(v int32) *CheckAuthCodeBindForExtResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetMessage(v string) *CheckAuthCodeBindForExtResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetRequestId(v string) *CheckAuthCodeBindForExtResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) SetResult(v *CheckAuthCodeBindForExtResponseBodyResult) *CheckAuthCodeBindForExtResponseBody {
	s.Result = v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckAuthCodeBindForExtResponseBodyResult struct {
	DeviceOpenInfo *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo `json:"DeviceOpenInfo,omitempty" xml:"DeviceOpenInfo,omitempty" type:"Struct"`
	UserOpenInfo   *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo   `json:"UserOpenInfo,omitempty" xml:"UserOpenInfo,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) GetDeviceOpenInfo() *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo {
	return s.DeviceOpenInfo
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) GetUserOpenInfo() *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo {
	return s.UserOpenInfo
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) SetDeviceOpenInfo(v *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) *CheckAuthCodeBindForExtResponseBodyResult {
	s.DeviceOpenInfo = v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) SetUserOpenInfo(v *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) *CheckAuthCodeBindForExtResponseBodyResult {
	s.UserOpenInfo = v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo struct {
	// example:
	//
	// A963*0158
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// DEVICE_ID
	//
	// example:
	//
	// DEVICE_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) GetId() *string {
	return s.Id
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) GetIdType() *string {
	return s.IdType
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) SetId(v string) *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) SetIdType(v string) *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo {
	s.IdType = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultDeviceOpenInfo) Validate() error {
	return dara.Validate(s)
}

type CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo struct {
	// example:
	//
	// 0963*0158
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// USER_ID
	//
	// example:
	//
	// USER_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
}

func (s CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) GetId() *string {
	return s.Id
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) GetIdType() *string {
	return s.IdType
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) SetId(v string) *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) SetIdType(v string) *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo {
	s.IdType = &v
	return s
}

func (s *CheckAuthCodeBindForExtResponseBodyResultUserOpenInfo) Validate() error {
	return dara.Validate(s)
}
