// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneBusinessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryPhoneBusinessProfileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryPhoneBusinessProfileResponseBody
	GetCode() *string
	SetData(v *QueryPhoneBusinessProfileResponseBodyData) *QueryPhoneBusinessProfileResponseBody
	GetData() *QueryPhoneBusinessProfileResponseBodyData
	SetMessage(v string) *QueryPhoneBusinessProfileResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPhoneBusinessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPhoneBusinessProfileResponseBody
	GetSuccess() *bool
}

type QueryPhoneBusinessProfileResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryPhoneBusinessProfileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
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

func (s QueryPhoneBusinessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryPhoneBusinessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPhoneBusinessProfileResponseBody) GetData() *QueryPhoneBusinessProfileResponseBodyData {
	return s.Data
}

func (s *QueryPhoneBusinessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPhoneBusinessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPhoneBusinessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPhoneBusinessProfileResponseBody) SetAccessDeniedDetail(v string) *QueryPhoneBusinessProfileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetCode(v string) *QueryPhoneBusinessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetData(v *QueryPhoneBusinessProfileResponseBodyData) *QueryPhoneBusinessProfileResponseBody {
	s.Data = v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetMessage(v string) *QueryPhoneBusinessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetRequestId(v string) *QueryPhoneBusinessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) SetSuccess(v bool) *QueryPhoneBusinessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPhoneBusinessProfileResponseBodyData struct {
	// Regarding.
	//
	// example:
	//
	// business profile
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// The address.
	//
	// example:
	//
	// Changsha
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The email address.
	//
	// example:
	//
	// aa@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The profile picture.
	//
	// example:
	//
	// https://....img
	ProfilePictureUrl *string `json:"ProfilePictureUrl,omitempty" xml:"ProfilePictureUrl,omitempty"`
	// The industry.
	//
	// example:
	//
	// Retail
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// The website.
	Websites []*string `json:"Websites,omitempty" xml:"Websites,omitempty" type:"Repeated"`
}

func (s QueryPhoneBusinessProfileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetAbout() *string {
	return s.About
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetProfilePictureUrl() *string {
	return s.ProfilePictureUrl
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetVertical() *string {
	return s.Vertical
}

func (s *QueryPhoneBusinessProfileResponseBodyData) GetWebsites() []*string {
	return s.Websites
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetAbout(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.About = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetAddress(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Address = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetDescription(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetEmail(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetProfilePictureUrl(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.ProfilePictureUrl = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetVertical(v string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Vertical = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) SetWebsites(v []*string) *QueryPhoneBusinessProfileResponseBodyData {
	s.Websites = v
	return s
}

func (s *QueryPhoneBusinessProfileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
