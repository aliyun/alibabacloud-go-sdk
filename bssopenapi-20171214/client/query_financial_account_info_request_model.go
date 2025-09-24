// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFinancialAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v int64) *QueryFinancialAccountInfoRequest
	GetUserId() *int64
}

type QueryFinancialAccountInfoRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1990699401005016
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFinancialAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFinancialAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryFinancialAccountInfoRequest) SetUserId(v int64) *QueryFinancialAccountInfoRequest {
	s.UserId = &v
	return s
}

func (s *QueryFinancialAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
