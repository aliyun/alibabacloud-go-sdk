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
	// This parameter is required.
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
