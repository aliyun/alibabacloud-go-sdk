// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetImageResponseBodyData) *GetImageResponseBody
	GetData() *GetImageResponseBodyData
	SetErrorCode(v string) *GetImageResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *GetImageResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GetImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageResponseBody
	GetSuccess() *bool
}

type GetImageResponseBody struct {
	// The details of the data.
	Data *GetImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25EA0A83-9007-4E87-808C-637BE1A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) GetData() *GetImageResponseBodyData {
	return s.Data
}

func (s *GetImageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetImageResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageResponseBody) SetData(v *GetImageResponseBodyData) *GetImageResponseBody {
	s.Data = v
	return s
}

func (s *GetImageResponseBody) SetErrorCode(v string) *GetImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetImageResponseBody) SetHttpCode(v string) *GetImageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetImageResponseBody) SetMessage(v string) *GetImageResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyData struct {
	// The full version number of the current instance image. The parameter is in the X.X.X.X format.
	//
	// example:
	//
	// 3.5.5.0
	CurrentVersionFullShowName *string `json:"CurrentVersionFullShowName,omitempty" xml:"CurrentVersionFullShowName,omitempty"`
	// The URL of the changelog for the maximum version to which the current version can be upgraded.
	//
	// example:
	//
	// https://xxxxx
	MaxVersionChangelogUrl *string `json:"MaxVersionChangelogUrl,omitempty" xml:"MaxVersionChangelogUrl,omitempty"`
	// The code of the maximum version to which the current version can be upgraded.
	//
	// example:
	//
	// ZooKeeper_3_6_3
	MaxVersionCode *string `json:"MaxVersionCode,omitempty" xml:"MaxVersionCode,omitempty"`
	// The full number of the maximum version to which the current version can be upgraded.
	//
	// example:
	//
	// 3.6.3.0
	MaxVersionFullShowName *string `json:"MaxVersionFullShowName,omitempty" xml:"MaxVersionFullShowName,omitempty"`
}

func (s GetImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyData) GetCurrentVersionFullShowName() *string {
	return s.CurrentVersionFullShowName
}

func (s *GetImageResponseBodyData) GetMaxVersionChangelogUrl() *string {
	return s.MaxVersionChangelogUrl
}

func (s *GetImageResponseBodyData) GetMaxVersionCode() *string {
	return s.MaxVersionCode
}

func (s *GetImageResponseBodyData) GetMaxVersionFullShowName() *string {
	return s.MaxVersionFullShowName
}

func (s *GetImageResponseBodyData) SetCurrentVersionFullShowName(v string) *GetImageResponseBodyData {
	s.CurrentVersionFullShowName = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionChangelogUrl(v string) *GetImageResponseBodyData {
	s.MaxVersionChangelogUrl = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionCode(v string) *GetImageResponseBodyData {
	s.MaxVersionCode = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionFullShowName(v string) *GetImageResponseBodyData {
	s.MaxVersionFullShowName = &v
	return s
}

func (s *GetImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
