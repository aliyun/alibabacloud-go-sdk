// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewOptimizeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSQLReviewQueryKey(v string) *GetSQLReviewOptimizeDetailRequest
	GetSQLReviewQueryKey() *string
	SetTid(v int64) *GetSQLReviewOptimizeDetailRequest
	GetTid() *int64
}

type GetSQLReviewOptimizeDetailRequest struct {
	// The key that is used to query the details of optimization suggestions. You can call the [ListSQLReviewOriginSQL](https://help.aliyun.com/document_detail/257870.html) operation to query the key.
	//
	// This parameter is required.
	//
	// example:
	//
	// a57e54ec5433475ea3082d882fdb****
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetSQLReviewOptimizeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailRequest) GetSQLReviewQueryKey() *string {
	return s.SQLReviewQueryKey
}

func (s *GetSQLReviewOptimizeDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetSQLReviewOptimizeDetailRequest) SetSQLReviewQueryKey(v string) *GetSQLReviewOptimizeDetailRequest {
	s.SQLReviewQueryKey = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailRequest) SetTid(v int64) *GetSQLReviewOptimizeDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailRequest) Validate() error {
	return dara.Validate(s)
}
