// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateSvcMapBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *BatchCreateSvcMapBindRequest
	GetJwtToken() *string
	SetMapIds(v []*int64) *BatchCreateSvcMapBindRequest
	GetMapIds() []*int64
	SetSvcId(v int64) *BatchCreateSvcMapBindRequest
	GetSvcId() *int64
}

type BatchCreateSvcMapBindRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	MapIds []*int64 `json:"MapIds,omitempty" xml:"MapIds,omitempty" type:"Repeated"`
	SvcId  *int64   `json:"SvcId,omitempty" xml:"SvcId,omitempty"`
}

func (s BatchCreateSvcMapBindRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateSvcMapBindRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateSvcMapBindRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *BatchCreateSvcMapBindRequest) GetMapIds() []*int64 {
	return s.MapIds
}

func (s *BatchCreateSvcMapBindRequest) GetSvcId() *int64 {
	return s.SvcId
}

func (s *BatchCreateSvcMapBindRequest) SetJwtToken(v string) *BatchCreateSvcMapBindRequest {
	s.JwtToken = &v
	return s
}

func (s *BatchCreateSvcMapBindRequest) SetMapIds(v []*int64) *BatchCreateSvcMapBindRequest {
	s.MapIds = v
	return s
}

func (s *BatchCreateSvcMapBindRequest) SetSvcId(v int64) *BatchCreateSvcMapBindRequest {
	s.SvcId = &v
	return s
}

func (s *BatchCreateSvcMapBindRequest) Validate() error {
	return dara.Validate(s)
}
