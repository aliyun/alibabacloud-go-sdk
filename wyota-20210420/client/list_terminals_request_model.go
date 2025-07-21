// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInManage(v bool) *ListTerminalsRequest
	GetInManage() *bool
	SetMainBizType(v string) *ListTerminalsRequest
	GetMainBizType() *string
	SetMaxResults(v int32) *ListTerminalsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTerminalsRequest
	GetNextToken() *string
	SetPasswordFreeLoginUser(v string) *ListTerminalsRequest
	GetPasswordFreeLoginUser() *string
	SetSearchKeyword(v string) *ListTerminalsRequest
	GetSearchKeyword() *string
	SetSerialNumbers(v []*string) *ListTerminalsRequest
	GetSerialNumbers() []*string
	SetTerminalGroupId(v string) *ListTerminalsRequest
	GetTerminalGroupId() *string
	SetUuids(v []*string) *ListTerminalsRequest
	GetUuids() []*string
	SetWithBindUser(v bool) *ListTerminalsRequest
	GetWithBindUser() *bool
}

type ListTerminalsRequest struct {
	InManage    *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	MainBizType *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6wQvfp7u1BJl4bxCAby41POSaYAlCvfULQpkAnb0ff****
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PasswordFreeLoginUser *string `json:"PasswordFreeLoginUser,omitempty" xml:"PasswordFreeLoginUser,omitempty"`
	// example:
	//
	// DemoDevice
	SearchKeyword *string   `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SerialNumbers []*string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// tg-default
	TerminalGroupId *string   `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuids           []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	WithBindUser    *bool     `json:"WithBindUser,omitempty" xml:"WithBindUser,omitempty"`
}

func (s ListTerminalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalsRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalsRequest) GetInManage() *bool {
	return s.InManage
}

func (s *ListTerminalsRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *ListTerminalsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTerminalsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerminalsRequest) GetPasswordFreeLoginUser() *string {
	return s.PasswordFreeLoginUser
}

func (s *ListTerminalsRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *ListTerminalsRequest) GetSerialNumbers() []*string {
	return s.SerialNumbers
}

func (s *ListTerminalsRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *ListTerminalsRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *ListTerminalsRequest) GetWithBindUser() *bool {
	return s.WithBindUser
}

func (s *ListTerminalsRequest) SetInManage(v bool) *ListTerminalsRequest {
	s.InManage = &v
	return s
}

func (s *ListTerminalsRequest) SetMainBizType(v string) *ListTerminalsRequest {
	s.MainBizType = &v
	return s
}

func (s *ListTerminalsRequest) SetMaxResults(v int32) *ListTerminalsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerminalsRequest) SetNextToken(v string) *ListTerminalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerminalsRequest) SetPasswordFreeLoginUser(v string) *ListTerminalsRequest {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *ListTerminalsRequest) SetSearchKeyword(v string) *ListTerminalsRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListTerminalsRequest) SetSerialNumbers(v []*string) *ListTerminalsRequest {
	s.SerialNumbers = v
	return s
}

func (s *ListTerminalsRequest) SetTerminalGroupId(v string) *ListTerminalsRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalsRequest) SetUuids(v []*string) *ListTerminalsRequest {
	s.Uuids = v
	return s
}

func (s *ListTerminalsRequest) SetWithBindUser(v bool) *ListTerminalsRequest {
	s.WithBindUser = &v
	return s
}

func (s *ListTerminalsRequest) Validate() error {
	return dara.Validate(s)
}
