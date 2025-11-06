// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetUserSettingResponseBody
	GetCode() *int32
	SetData(v *GetUserSettingResponseBodyData) *GetUserSettingResponseBody
	GetData() *GetUserSettingResponseBodyData
	SetMessage(v string) *GetUserSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserSettingResponseBody
	GetSuccess() *bool
}

type GetUserSettingResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetUserSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetUserSettingResponseBody) GetData() *GetUserSettingResponseBodyData {
	return s.Data
}

func (s *GetUserSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserSettingResponseBody) SetCode(v int32) *GetUserSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserSettingResponseBody) SetData(v *GetUserSettingResponseBodyData) *GetUserSettingResponseBody {
	s.Data = v
	return s
}

func (s *GetUserSettingResponseBody) SetMessage(v string) *GetUserSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserSettingResponseBody) SetRequestId(v string) *GetUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserSettingResponseBody) SetSuccess(v bool) *GetUserSettingResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserSettingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserSettingResponseBodyData struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserSettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserSettingResponseBodyData) GetLogstore() *string {
	return s.Logstore
}

func (s *GetUserSettingResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserSettingResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *GetUserSettingResponseBodyData) SetLogstore(v string) *GetUserSettingResponseBodyData {
	s.Logstore = &v
	return s
}

func (s *GetUserSettingResponseBodyData) SetRegionId(v string) *GetUserSettingResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUserSettingResponseBodyData) SetUserId(v int64) *GetUserSettingResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetUserSettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
