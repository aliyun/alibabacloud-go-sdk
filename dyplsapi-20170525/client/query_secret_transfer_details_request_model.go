// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretTransferDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *QuerySecretTransferDetailsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QuerySecretTransferDetailsRequest
	GetPageSize() *int64
	SetRecordId(v int64) *QuerySecretTransferDetailsRequest
	GetRecordId() *int64
}

type QuerySecretTransferDetailsRequest struct {
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 迁移记录的id
	//
	// This parameter is required.
	//
	// example:
	//
	// 55
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s QuerySecretTransferDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretTransferDetailsRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretTransferDetailsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QuerySecretTransferDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QuerySecretTransferDetailsRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *QuerySecretTransferDetailsRequest) SetPageNo(v int64) *QuerySecretTransferDetailsRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySecretTransferDetailsRequest) SetPageSize(v int64) *QuerySecretTransferDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySecretTransferDetailsRequest) SetRecordId(v int64) *QuerySecretTransferDetailsRequest {
	s.RecordId = &v
	return s
}

func (s *QuerySecretTransferDetailsRequest) Validate() error {
	return dara.Validate(s)
}
