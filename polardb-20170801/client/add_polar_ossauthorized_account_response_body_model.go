// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarOSSAuthorizedAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserArnIds(v string) *AddPolarOSSAuthorizedAccountResponseBody
	GetAuthorizedUserArnIds() *string
	SetAuthorizedUserIds(v string) *AddPolarOSSAuthorizedAccountResponseBody
	GetAuthorizedUserIds() *string
	SetPfsInstanceId(v string) *AddPolarOSSAuthorizedAccountResponseBody
	GetPfsInstanceId() *string
	SetRequestId(v string) *AddPolarOSSAuthorizedAccountResponseBody
	GetRequestId() *string
}

type AddPolarOSSAuthorizedAccountResponseBody struct {
	// The updated list of RAM role ARNs, separated by commas.
	//
	// example:
	//
	// arn:sts::123456:assumed-role/myrole/*
	AuthorizedUserArnIds *string `json:"AuthorizedUserArnIds,omitempty" xml:"AuthorizedUserArnIds,omitempty"`
	// The updated list of UIDs, separated by commas.
	//
	// example:
	//
	// 1234567890,9876543210
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	// The cold storage instance ID.
	//
	// example:
	//
	// pfs-xxxxxxxxxxxxxxxxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPolarOSSAuthorizedAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPolarOSSAuthorizedAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) GetAuthorizedUserArnIds() *string {
	return s.AuthorizedUserArnIds
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) SetAuthorizedUserArnIds(v string) *AddPolarOSSAuthorizedAccountResponseBody {
	s.AuthorizedUserArnIds = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) SetAuthorizedUserIds(v string) *AddPolarOSSAuthorizedAccountResponseBody {
	s.AuthorizedUserIds = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) SetPfsInstanceId(v string) *AddPolarOSSAuthorizedAccountResponseBody {
	s.PfsInstanceId = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) SetRequestId(v string) *AddPolarOSSAuthorizedAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
