// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteTemplateRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteTemplateRequest
	GetOwnerId() *int64
}

type DeleteTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTemplateRequest) SetId(v string) *DeleteTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteTemplateRequest) SetOwnerId(v int64) *DeleteTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
