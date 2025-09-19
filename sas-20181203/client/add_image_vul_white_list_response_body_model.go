// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageVulWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddImageVulWhiteListResponseBody
	GetCode() *string
	SetData(v bool) *AddImageVulWhiteListResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddImageVulWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddImageVulWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddImageVulWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddImageVulWhiteListResponseBody
	GetSuccess() *bool
}

type AddImageVulWhiteListResponseBody struct {
	// The status code returned. A value of **200*	- indicates that the request was successful. Other values indicate that the request failed. You can identify the cause of the failure based on the value of this parameter.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the image vulnerability is added to the whitelist. Valid values:
	//
	// 	- **true**: The image vulnerability is added to the whitelist.
	//
	// 	- **false**: The image vulnerability is not added to the whitelist.
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
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageVulWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageVulWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageVulWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddImageVulWhiteListResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddImageVulWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddImageVulWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddImageVulWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageVulWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImageVulWhiteListResponseBody) SetCode(v string) *AddImageVulWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) SetData(v bool) *AddImageVulWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) SetHttpStatusCode(v int32) *AddImageVulWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) SetMessage(v string) *AddImageVulWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) SetRequestId(v string) *AddImageVulWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) SetSuccess(v bool) *AddImageVulWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *AddImageVulWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
