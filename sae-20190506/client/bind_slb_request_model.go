// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BindSlbRequest
	GetAppId() *string
	SetInternet(v string) *BindSlbRequest
	GetInternet() *string
	SetInternetSlbChargeType(v string) *BindSlbRequest
	GetInternetSlbChargeType() *string
	SetInternetSlbId(v string) *BindSlbRequest
	GetInternetSlbId() *string
	SetIntranet(v string) *BindSlbRequest
	GetIntranet() *string
	SetIntranetSlbChargeType(v string) *BindSlbRequest
	GetIntranetSlbChargeType() *string
	SetIntranetSlbId(v string) *BindSlbRequest
	GetIntranetSlbId() *string
}

type BindSlbRequest struct {
	// 0099b7be-5f5b-4512-a7fc-56049ef1\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// [{"port":80,"targetPort":8080,"protocol":"TCP"}]
	//
	// example:
	//
	// [{"port":80,"targetPort":8080,"protocol":"TCP"}]
	Internet *string `json:"Internet,omitempty" xml:"Internet,omitempty"`
	// The billing method of an Internet-facing SLB instance. The following billing methods are supported:
	//
	// 	- **PayBySpec**: Pay-by-specification.
	//
	// 	- **PayByCLCU**: Pay-by-CLCU.
	//
	// example:
	//
	// PayBySpec
	InternetSlbChargeType *string `json:"InternetSlbChargeType,omitempty" xml:"InternetSlbChargeType,omitempty"`
	// lb-bp1tg0k6d9nqaw7l1\\*\\*\\*\\*
	//
	// example:
	//
	// lb-bp1tg0k6d9nqaw7l1****
	InternetSlbId *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	// [{"port":80,"targetPort":8080,"protocol":"TCP"}]
	//
	// example:
	//
	// [{"port":80,"targetPort":8080,"protocol":"TCP"}]
	Intranet *string `json:"Intranet,omitempty" xml:"Intranet,omitempty"`
	// The billing method of an Internal-facing SLB instance. The following billing methods are supported:
	//
	// 	- **PayBySpec**: Pay-by-specification.
	//
	// 	- **PayByCLCU**: Pay-by-CLCU.
	//
	// example:
	//
	// PayBySpec
	IntranetSlbChargeType *string `json:"IntranetSlbChargeType,omitempty" xml:"IntranetSlbChargeType,omitempty"`
	// lb-bp1tg0k6d9nqaw7l1\\*\\*\\*\\*
	//
	// example:
	//
	// lb-bp1tg0k6d9nqaw7l1****
	IntranetSlbId *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
}

func (s BindSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSlbRequest) GoString() string {
	return s.String()
}

func (s *BindSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindSlbRequest) GetInternet() *string {
	return s.Internet
}

func (s *BindSlbRequest) GetInternetSlbChargeType() *string {
	return s.InternetSlbChargeType
}

func (s *BindSlbRequest) GetInternetSlbId() *string {
	return s.InternetSlbId
}

func (s *BindSlbRequest) GetIntranet() *string {
	return s.Intranet
}

func (s *BindSlbRequest) GetIntranetSlbChargeType() *string {
	return s.IntranetSlbChargeType
}

func (s *BindSlbRequest) GetIntranetSlbId() *string {
	return s.IntranetSlbId
}

func (s *BindSlbRequest) SetAppId(v string) *BindSlbRequest {
	s.AppId = &v
	return s
}

func (s *BindSlbRequest) SetInternet(v string) *BindSlbRequest {
	s.Internet = &v
	return s
}

func (s *BindSlbRequest) SetInternetSlbChargeType(v string) *BindSlbRequest {
	s.InternetSlbChargeType = &v
	return s
}

func (s *BindSlbRequest) SetInternetSlbId(v string) *BindSlbRequest {
	s.InternetSlbId = &v
	return s
}

func (s *BindSlbRequest) SetIntranet(v string) *BindSlbRequest {
	s.Intranet = &v
	return s
}

func (s *BindSlbRequest) SetIntranetSlbChargeType(v string) *BindSlbRequest {
	s.IntranetSlbChargeType = &v
	return s
}

func (s *BindSlbRequest) SetIntranetSlbId(v string) *BindSlbRequest {
	s.IntranetSlbId = &v
	return s
}

func (s *BindSlbRequest) Validate() error {
	return dara.Validate(s)
}
