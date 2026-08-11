// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlashSmsProvidersResponseBody
	GetCode() *string
	SetData(v []*ListFlashSmsProvidersResponseBodyData) *ListFlashSmsProvidersResponseBody
	GetData() []*ListFlashSmsProvidersResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsProvidersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsProvidersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListFlashSmsProvidersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListFlashSmsProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFlashSmsProvidersResponseBody
	GetSuccess() *bool
}

type ListFlashSmsProvidersResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListFlashSmsProvidersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=outb003
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFlashSmsProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlashSmsProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlashSmsProvidersResponseBody) GetData() []*ListFlashSmsProvidersResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlashSmsProvidersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListFlashSmsProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlashSmsProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFlashSmsProvidersResponseBody) SetCode(v string) *ListFlashSmsProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetData(v []*ListFlashSmsProvidersResponseBodyData) *ListFlashSmsProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetHttpStatusCode(v int32) *ListFlashSmsProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetMessage(v string) *ListFlashSmsProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetParams(v []*string) *ListFlashSmsProvidersResponseBody {
	s.Params = v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetRequestId(v string) *ListFlashSmsProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) SetSuccess(v bool) *ListFlashSmsProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBody) Validate() error {
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

type ListFlashSmsProvidersResponseBodyData struct {
	// The provider parameter information.
	//
	// example:
	//
	// {\\"Pwd\\":{\\"displayName\\":\\"密码\\",\\"dataType\\":\\"TEXT\\",\\"displayOrder\\":2,\\"required\\":true,\\"maxLength\\":64},\\"User\\":{\\"displayName\\":\\"用户\\",\\"dataType\\":\\"TEXT\\",\\"displayOrder\\":1,\\"required\\":true,\\"maxLength\\":64},\\"Account\\":{\\"displayName\\":\\"账号\\",\\"dataType\\":\\"TEXT\\",\\"displayOrder\\":3,\\"required\\":true,\\"maxLength\\":64}}
	ProfileSchema *string `json:"ProfileSchema,omitempty" xml:"ProfileSchema,omitempty"`
	// The provider ID.
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// The provider name.
	//
	// example:
	//
	// 北京优音通信有限公司
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s ListFlashSmsProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsProvidersResponseBodyData) GetProfileSchema() *string {
	return s.ProfileSchema
}

func (s *ListFlashSmsProvidersResponseBodyData) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsProvidersResponseBodyData) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListFlashSmsProvidersResponseBodyData) SetProfileSchema(v string) *ListFlashSmsProvidersResponseBodyData {
	s.ProfileSchema = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBodyData) SetProviderId(v string) *ListFlashSmsProvidersResponseBodyData {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBodyData) SetProviderName(v string) *ListFlashSmsProvidersResponseBodyData {
	s.ProviderName = &v
	return s
}

func (s *ListFlashSmsProvidersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
