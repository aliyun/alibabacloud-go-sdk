// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountContactQueryPageListRequest
	GetAppName() *string
	SetOrientedEcId(v string) *AccountContactQueryPageListRequest
	GetOrientedEcId() *string
	SetOrientedLeId(v string) *AccountContactQueryPageListRequest
	GetOrientedLeId() *string
	SetOrientedNbId(v string) *AccountContactQueryPageListRequest
	GetOrientedNbId() *string
	SetPageNo(v int32) *AccountContactQueryPageListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *AccountContactQueryPageListRequest
	GetPageSize() *int32
	SetPrivateContact(v bool) *AccountContactQueryPageListRequest
	GetPrivateContact() *bool
	SetQuery(v string) *AccountContactQueryPageListRequest
	GetQuery() *string
	SetSharedContact(v bool) *AccountContactQueryPageListRequest
	GetSharedContact() *bool
	SetShowCompleteInfo(v bool) *AccountContactQueryPageListRequest
	GetShowCompleteInfo() *bool
}

type AccountContactQueryPageListRequest struct {
	// example:
	//
	// xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// null
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	// example:
	//
	// null
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	// example:
	//
	// null
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PrivateContact *bool `json:"PrivateContact,omitempty" xml:"PrivateContact,omitempty"`
	// example:
	//
	// keyword
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// false
	SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
	// example:
	//
	// false
	ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
}

func (s AccountContactQueryPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryPageListRequest) GoString() string {
	return s.String()
}

func (s *AccountContactQueryPageListRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountContactQueryPageListRequest) GetOrientedEcId() *string {
	return s.OrientedEcId
}

func (s *AccountContactQueryPageListRequest) GetOrientedLeId() *string {
	return s.OrientedLeId
}

func (s *AccountContactQueryPageListRequest) GetOrientedNbId() *string {
	return s.OrientedNbId
}

func (s *AccountContactQueryPageListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *AccountContactQueryPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AccountContactQueryPageListRequest) GetPrivateContact() *bool {
	return s.PrivateContact
}

func (s *AccountContactQueryPageListRequest) GetQuery() *string {
	return s.Query
}

func (s *AccountContactQueryPageListRequest) GetSharedContact() *bool {
	return s.SharedContact
}

func (s *AccountContactQueryPageListRequest) GetShowCompleteInfo() *bool {
	return s.ShowCompleteInfo
}

func (s *AccountContactQueryPageListRequest) SetAppName(v string) *AccountContactQueryPageListRequest {
	s.AppName = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetOrientedEcId(v string) *AccountContactQueryPageListRequest {
	s.OrientedEcId = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetOrientedLeId(v string) *AccountContactQueryPageListRequest {
	s.OrientedLeId = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetOrientedNbId(v string) *AccountContactQueryPageListRequest {
	s.OrientedNbId = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetPageNo(v int32) *AccountContactQueryPageListRequest {
	s.PageNo = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetPageSize(v int32) *AccountContactQueryPageListRequest {
	s.PageSize = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetPrivateContact(v bool) *AccountContactQueryPageListRequest {
	s.PrivateContact = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetQuery(v string) *AccountContactQueryPageListRequest {
	s.Query = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetSharedContact(v bool) *AccountContactQueryPageListRequest {
	s.SharedContact = &v
	return s
}

func (s *AccountContactQueryPageListRequest) SetShowCompleteInfo(v bool) *AccountContactQueryPageListRequest {
	s.ShowCompleteInfo = &v
	return s
}

func (s *AccountContactQueryPageListRequest) Validate() error {
	return dara.Validate(s)
}
