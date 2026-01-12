// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSvcMapBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *BatchDeleteSvcMapBindRequest
	GetIds() []*int64
	SetJwtToken(v string) *BatchDeleteSvcMapBindRequest
	GetJwtToken() *string
}

type BatchDeleteSvcMapBindRequest struct {
	// This parameter is required.
	Ids      []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	JwtToken *string  `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s BatchDeleteSvcMapBindRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSvcMapBindRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteSvcMapBindRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *BatchDeleteSvcMapBindRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *BatchDeleteSvcMapBindRequest) SetIds(v []*int64) *BatchDeleteSvcMapBindRequest {
	s.Ids = v
	return s
}

func (s *BatchDeleteSvcMapBindRequest) SetJwtToken(v string) *BatchDeleteSvcMapBindRequest {
	s.JwtToken = &v
	return s
}

func (s *BatchDeleteSvcMapBindRequest) Validate() error {
	return dara.Validate(s)
}
