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
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The call is successful.
	//
	// - **3xx**: The call is redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The version information.
	Data []*ListAppVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information about the call.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the historical versions of the application were successfully queried. Valid values:
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
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
	// The download URL of the code package. If you uploaded the package using SAE, note the following:
	//
	// - This URL is not a direct download link. You must call the GetPackageVersionAccessableUrl operation to obtain a downloadable URL that is valid for 10 minutes.
	//
	// - SAE stores the package for a maximum of 90 days. After this period, the URL is not returned and the package cannot be downloaded.
	BuildPackageUrl *string `json:"BuildPackageUrl,omitempty" xml:"BuildPackageUrl,omitempty"`
	// The time when the version was created.
	//
	// example:
	//
	// 1590124643553
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version ID.
	//
	// example:
	//
	// a0ce266c-d354-423a-9bd6-4083405a****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The application type. Valid values:
	//
	// - **image**: The application is deployed using an image.
	//
	// - **url**: The application is deployed using a code package.
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the WAR package.
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
