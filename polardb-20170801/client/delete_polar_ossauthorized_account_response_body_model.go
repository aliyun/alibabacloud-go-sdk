// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarOSSAuthorizedAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserArnIds(v string) *DeletePolarOSSAuthorizedAccountResponseBody
	GetAuthorizedUserArnIds() *string
	SetAuthorizedUserIds(v string) *DeletePolarOSSAuthorizedAccountResponseBody
	GetAuthorizedUserIds() *string
	SetPfsInstanceId(v string) *DeletePolarOSSAuthorizedAccountResponseBody
	GetPfsInstanceId() *string
	SetRequestId(v string) *DeletePolarOSSAuthorizedAccountResponseBody
	GetRequestId() *string
}

type DeletePolarOSSAuthorizedAccountResponseBody struct {
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

func (s DeletePolarOSSAuthorizedAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarOSSAuthorizedAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) GetAuthorizedUserArnIds() *string {
	return s.AuthorizedUserArnIds
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) SetAuthorizedUserArnIds(v string) *DeletePolarOSSAuthorizedAccountResponseBody {
	s.AuthorizedUserArnIds = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) SetAuthorizedUserIds(v string) *DeletePolarOSSAuthorizedAccountResponseBody {
	s.AuthorizedUserIds = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) SetPfsInstanceId(v string) *DeletePolarOSSAuthorizedAccountResponseBody {
	s.PfsInstanceId = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) SetRequestId(v string) *DeletePolarOSSAuthorizedAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
