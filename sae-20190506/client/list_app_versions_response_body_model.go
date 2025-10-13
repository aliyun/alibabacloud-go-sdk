// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppVersionsResponseBody
	GetCode() *string
	SetData(v []*ListAppVersionsResponseBodyData) *ListAppVersionsResponseBody
	GetData() []*ListAppVersionsResponseBodyData
	SetErrorCode(v string) *ListAppVersionsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppVersionsResponseBody
	GetSuccess() *bool
}

type ListAppVersionsResponseBody struct {
	// Indicates whether the historical versions of the application were obtained. Valid values:
	//
	// 	- **true**: indicates that the historical versions of the application were obtained.
	//
	// 	- **false**: indicates that the historical versions of the application could not be obtained.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the versions.
	Data []*ListAppVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the versions.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppVersionsResponseBody) GetData() []*ListAppVersionsResponseBodyData {
	return s.Data
}

func (s *ListAppVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppVersionsResponseBody) SetCode(v string) *ListAppVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetData(v []*ListAppVersionsResponseBodyData) *ListAppVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppVersionsResponseBody) SetErrorCode(v string) *ListAppVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetMessage(v string) *ListAppVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetRequestId(v string) *ListAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetSuccess(v bool) *ListAppVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppVersionsResponseBody) Validate() error {
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

type ListAppVersionsResponseBodyData struct {
	// The URL of the code package. If you use the SAE console to upload the code package, take note of the following items:
	//
	// 	- You cannot download the URL. You must call the GetPackageVersionAccessableUrl operation to obtain the URL. The obtained URL is valid for 10 minutes.
	//
	// 	- SAE can retain the package up to 90 days. After 90 days, the URL cannot be returned or downloaded.
	BuildPackageUrl *string `json:"BuildPackageUrl,omitempty" xml:"BuildPackageUrl,omitempty"`
	// The download link of the WAR or JAR package. This parameter is returned when the **Type*	- parameter is set to **url**.
	//
	// example:
	//
	// 1590124643553
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// a0ce266c-d354-423a-9bd6-4083405a****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The deployment method of the application. Valid values:
	//
	// 	- **image**: indicates that the application is deployed by using an image.
	//
	// 	- **url**: indicates that the application is deployed by using a code package.
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the image.
	WarUrl *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
}

func (s ListAppVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponseBodyData) GetBuildPackageUrl() *string {
	return s.BuildPackageUrl
}

func (s *ListAppVersionsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAppVersionsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListAppVersionsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListAppVersionsResponseBodyData) GetWarUrl() *string {
	return s.WarUrl
}

func (s *ListAppVersionsResponseBodyData) SetBuildPackageUrl(v string) *ListAppVersionsResponseBodyData {
	s.BuildPackageUrl = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetCreateTime(v string) *ListAppVersionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetId(v string) *ListAppVersionsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetType(v string) *ListAppVersionsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetWarUrl(v string) *ListAppVersionsResponseBodyData {
	s.WarUrl = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
