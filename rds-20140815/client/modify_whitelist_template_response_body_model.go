// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyWhitelistTemplateResponseBody
	GetCode() *string
	SetData(v *ModifyWhitelistTemplateResponseBodyData) *ModifyWhitelistTemplateResponseBody
	GetData() *ModifyWhitelistTemplateResponseBodyData
	SetHttpStatusCode(v int32) *ModifyWhitelistTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyWhitelistTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyWhitelistTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyWhitelistTemplateResponseBody
	GetSuccess() *bool
}

type ModifyWhitelistTemplateResponseBody struct {
	// The response code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **401**: identity authentication failed
	//
	// 	- **404**: request page not found
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ModifyWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyWhitelistTemplateResponseBody) GetData() *ModifyWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *ModifyWhitelistTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyWhitelistTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWhitelistTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyWhitelistTemplateResponseBody) SetCode(v string) *ModifyWhitelistTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) SetData(v *ModifyWhitelistTemplateResponseBodyData) *ModifyWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) SetHttpStatusCode(v int32) *ModifyWhitelistTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) SetMessage(v string) *ModifyWhitelistTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) SetRequestId(v string) *ModifyWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) SetSuccess(v bool) *ModifyWhitelistTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyWhitelistTemplateResponseBodyData struct {
	// The status code returned. Valid values:
	//
	// 	- **ok**: The request is successful.
	//
	// 	- **error**: The request fails.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistTemplateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModifyWhitelistTemplateResponseBodyData) SetStatus(v string) *ModifyWhitelistTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModifyWhitelistTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
