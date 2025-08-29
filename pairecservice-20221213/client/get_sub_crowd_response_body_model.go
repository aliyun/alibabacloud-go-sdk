// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *GetSubCrowdResponseBody
	GetGmtCreateTime() *string
	SetQuantity(v string) *GetSubCrowdResponseBody
	GetQuantity() *string
	SetRequestId(v string) *GetSubCrowdResponseBody
	GetRequestId() *string
	SetSource(v string) *GetSubCrowdResponseBody
	GetSource() *string
	SetUsers(v string) *GetSubCrowdResponseBody
	GetUsers() *string
}

type GetSubCrowdResponseBody struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 3
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s GetSubCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubCrowdResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetSubCrowdResponseBody) GetQuantity() *string {
	return s.Quantity
}

func (s *GetSubCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubCrowdResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetSubCrowdResponseBody) GetUsers() *string {
	return s.Users
}

func (s *GetSubCrowdResponseBody) SetGmtCreateTime(v string) *GetSubCrowdResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetQuantity(v string) *GetSubCrowdResponseBody {
	s.Quantity = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetRequestId(v string) *GetSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetSource(v string) *GetSubCrowdResponseBody {
	s.Source = &v
	return s
}

func (s *GetSubCrowdResponseBody) SetUsers(v string) *GetSubCrowdResponseBody {
	s.Users = &v
	return s
}

func (s *GetSubCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
