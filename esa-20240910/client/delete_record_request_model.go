// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *DeleteRecordRequest
	GetRecordId() *int64
	SetSecurityToken(v string) *DeleteRecordRequest
	GetSecurityToken() *string
}

type DeleteRecordRequest struct {
	// The ID of the DNS record. You can get this ID by calling the [ListRecords](https://help.aliyun.com/document_detail/2850265.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId      *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteRecordRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteRecordRequest) SetRecordId(v int64) *DeleteRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteRecordRequest) SetSecurityToken(v string) *DeleteRecordRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteRecordRequest) Validate() error {
	return dara.Validate(s)
}
