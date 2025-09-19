// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageVulWhitelistTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateImageVulWhitelistTargetResponseBody
	GetCode() *string
	SetData(v bool) *UpdateImageVulWhitelistTargetResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateImageVulWhitelistTargetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateImageVulWhitelistTargetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageVulWhitelistTargetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageVulWhitelistTargetResponseBody
	GetSuccess() *bool
}

type UpdateImageVulWhitelistTargetResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the image vulnerability is added to the whitelist. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
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
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s UpdateImageVulWhitelistTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageVulWhitelistTargetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageVulWhitelistTargetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetCode(v string) *UpdateImageVulWhitelistTargetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetData(v bool) *UpdateImageVulWhitelistTargetResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetHttpStatusCode(v int32) *UpdateImageVulWhitelistTargetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetMessage(v string) *UpdateImageVulWhitelistTargetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetRequestId(v string) *UpdateImageVulWhitelistTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) SetSuccess(v bool) *UpdateImageVulWhitelistTargetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
