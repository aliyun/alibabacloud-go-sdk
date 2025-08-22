// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAppResponseBody
	GetCode() *int32
	SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody
	GetData() *CreateAppResponseBodyData
	SetMessage(v string) *CreateAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAppResponseBody
	GetSuccess() *bool
}

type CreateAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAppResponseBody) GetData() *CreateAppResponseBodyData {
	return s.Data
}

func (s *CreateAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAppResponseBody) SetCode(v int32) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetSuccess(v bool) *CreateAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyData struct {
	// example:
	//
	// 4a0fae835cd741f3b12376d8a5a8e549v3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 10
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
}

func (s CreateAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateAppResponseBodyData) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *CreateAppResponseBodyData) SetAccessToken(v string) *CreateAppResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *CreateAppResponseBodyData) SetAppGroupId(v int64) *CreateAppResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *CreateAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
