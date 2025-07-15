// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceOrderConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataSourceOrderConfigResponseBody
	GetCode() *string
	SetData(v *GetDataSourceOrderConfigResponseBodyData) *GetDataSourceOrderConfigResponseBody
	GetData() *GetDataSourceOrderConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetDataSourceOrderConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataSourceOrderConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataSourceOrderConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataSourceOrderConfigResponseBody
	GetSuccess() *bool
}

type GetDataSourceOrderConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataSourceOrderConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataSourceOrderConfigResponseBody) GetData() *GetDataSourceOrderConfigResponseBodyData {
	return s.Data
}

func (s *GetDataSourceOrderConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataSourceOrderConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataSourceOrderConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceOrderConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataSourceOrderConfigResponseBody) SetCode(v string) *GetDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetData(v *GetDataSourceOrderConfigResponseBodyData) *GetDataSourceOrderConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *GetDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetMessage(v string) *GetDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetRequestId(v string) *GetDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetSuccess(v bool) *GetDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceOrderConfigResponseBodyData struct {
	// example:
	//
	// 1
	TotalDocSize             *int32                                                              `json:"TotalDocSize,omitempty" xml:"TotalDocSize,omitempty"`
	UserConfigDataSourceList []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s GetDataSourceOrderConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyData) GetTotalDocSize() *int32 {
	return s.TotalDocSize
}

func (s *GetDataSourceOrderConfigResponseBodyData) GetUserConfigDataSourceList() []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	return s.UserConfigDataSourceList
}

func (s *GetDataSourceOrderConfigResponseBodyData) SetTotalDocSize(v int32) *GetDataSourceOrderConfigResponseBodyData {
	s.TotalDocSize = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyData) SetUserConfigDataSourceList(v []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) *GetDataSourceOrderConfigResponseBodyData {
	s.UserConfigDataSourceList = v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList struct {
	// example:
	//
	// QuarkCommonNews
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 20
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// SystemSearch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GetCode() *string {
	return s.Code
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GetEnable() *bool {
	return s.Enable
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GetName() *string {
	return s.Name
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GetNumber() *int32 {
	return s.Number
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GetType() *string {
	return s.Type
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetCode(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetEnable(v bool) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Enable = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetName(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetNumber(v int32) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetType(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Type = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) Validate() error {
	return dara.Validate(s)
}
