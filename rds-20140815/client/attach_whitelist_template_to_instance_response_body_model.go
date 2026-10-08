// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachWhitelistTemplateToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttachWhitelistTemplateToInstanceResponseBody
	GetCode() *string
	SetData(v *AttachWhitelistTemplateToInstanceResponseBodyData) *AttachWhitelistTemplateToInstanceResponseBody
	GetData() *AttachWhitelistTemplateToInstanceResponseBodyData
	SetHttpStatusCode(v int32) *AttachWhitelistTemplateToInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AttachWhitelistTemplateToInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachWhitelistTemplateToInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachWhitelistTemplateToInstanceResponseBody
	GetSuccess() *bool
}

type AttachWhitelistTemplateToInstanceResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: Normal.
	//
	// - **400**: Client fault.
	//
	// - **401**: Authentication failed.
	//
	// - **404**: Request page not found.
	//
	// - **500**: Server fault.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *AttachWhitelistTemplateToInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// - **200**: Success.
	//
	// - **400**: Client error.
	//
	// - **500**: Server error.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetData() *AttachWhitelistTemplateToInstanceResponseBodyData {
	return s.Data
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetCode(v string) *AttachWhitelistTemplateToInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetData(v *AttachWhitelistTemplateToInstanceResponseBodyData) *AttachWhitelistTemplateToInstanceResponseBody {
	s.Data = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetHttpStatusCode(v int32) *AttachWhitelistTemplateToInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetMessage(v string) *AttachWhitelistTemplateToInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetRequestId(v string) *AttachWhitelistTemplateToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetSuccess(v bool) *AttachWhitelistTemplateToInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachWhitelistTemplateToInstanceResponseBodyData struct {
	// The return status. Valid values:
	//
	// - **ok**: Success.
	//
	// - **error**: Error.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) SetStatus(v string) *AttachWhitelistTemplateToInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
