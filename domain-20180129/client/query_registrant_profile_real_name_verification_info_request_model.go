// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfileRealNameVerificationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFetchImage(v bool) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetFetchImage() *bool
	SetLang(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetUserClientIp() *string
}

type QueryRegistrantProfileRealNameVerificationInfoRequest struct {
	// Specifies whether to retrieve the identity verification image. Valid values:
	//
	// - **true**: Retrieve the image.
	//
	// - **false**: Do not retrieve the image.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	FetchImage *bool `json:"FetchImage,omitempty" xml:"FetchImage,omitempty"`
	// The language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the information template to be queried.
	//
	// The system automatically generates this ID after the information template is created. You can call the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) API to query the information template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// The user IP address. You can set it to 127.0.0.1.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetFetchImage() *bool {
	return s.FetchImage
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetFetchImage(v bool) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.FetchImage = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetLang(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetUserClientIp(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) Validate() error {
	return dara.Validate(s)
}
