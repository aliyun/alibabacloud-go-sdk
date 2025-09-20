// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeParseSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeParseSettingResponseBody
	GetCode() *string
	SetData(v *ChangeParseSettingResponseBodyData) *ChangeParseSettingResponseBody
	GetData() *ChangeParseSettingResponseBodyData
	SetMessage(v string) *ChangeParseSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeParseSettingResponseBody
	GetRequestId() *string
	SetStatus(v string) *ChangeParseSettingResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ChangeParseSettingResponseBody
	GetSuccess() *bool
}

type ChangeParseSettingResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ChangeParseSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BA8ADD9-53D6-53F0-918F-A1E776AD230E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeParseSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeParseSettingResponseBody) GetData() *ChangeParseSettingResponseBodyData {
	return s.Data
}

func (s *ChangeParseSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeParseSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeParseSettingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ChangeParseSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeParseSettingResponseBody) SetCode(v string) *ChangeParseSettingResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeParseSettingResponseBody) SetData(v *ChangeParseSettingResponseBodyData) *ChangeParseSettingResponseBody {
	s.Data = v
	return s
}

func (s *ChangeParseSettingResponseBody) SetMessage(v string) *ChangeParseSettingResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeParseSettingResponseBody) SetRequestId(v string) *ChangeParseSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeParseSettingResponseBody) SetStatus(v string) *ChangeParseSettingResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeParseSettingResponseBody) SetSuccess(v bool) *ChangeParseSettingResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeParseSettingResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeParseSettingResponseBodyData struct {
	// example:
	//
	// true
	ChangeResult *bool `json:"ChangeResult,omitempty" xml:"ChangeResult,omitempty"`
}

func (s ChangeParseSettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeParseSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeParseSettingResponseBodyData) GetChangeResult() *bool {
	return s.ChangeResult
}

func (s *ChangeParseSettingResponseBodyData) SetChangeResult(v bool) *ChangeParseSettingResponseBodyData {
	s.ChangeResult = &v
	return s
}

func (s *ChangeParseSettingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
