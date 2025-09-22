// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleActivityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *QuerySingleActivityInfoRequest
	GetActivityId() *string
	SetCompanyName(v string) *QuerySingleActivityInfoRequest
	GetCompanyName() *string
	SetCustomerName(v string) *QuerySingleActivityInfoRequest
	GetCustomerName() *string
	SetMobile(v string) *QuerySingleActivityInfoRequest
	GetMobile() *string
	SetQRCode(v string) *QuerySingleActivityInfoRequest
	GetQRCode() *string
}

type QuerySingleActivityInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ActivityId   *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	CompanyName  *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// 12233445
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	QRCode *string `json:"QRCode,omitempty" xml:"QRCode,omitempty"`
}

func (s QuerySingleActivityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleActivityInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *QuerySingleActivityInfoRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QuerySingleActivityInfoRequest) GetCustomerName() *string {
	return s.CustomerName
}

func (s *QuerySingleActivityInfoRequest) GetMobile() *string {
	return s.Mobile
}

func (s *QuerySingleActivityInfoRequest) GetQRCode() *string {
	return s.QRCode
}

func (s *QuerySingleActivityInfoRequest) SetActivityId(v string) *QuerySingleActivityInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetCompanyName(v string) *QuerySingleActivityInfoRequest {
	s.CompanyName = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetCustomerName(v string) *QuerySingleActivityInfoRequest {
	s.CustomerName = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetMobile(v string) *QuerySingleActivityInfoRequest {
	s.Mobile = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) SetQRCode(v string) *QuerySingleActivityInfoRequest {
	s.QRCode = &v
	return s
}

func (s *QuerySingleActivityInfoRequest) Validate() error {
	return dara.Validate(s)
}
