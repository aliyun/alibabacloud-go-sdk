// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDcdnWafGroupRequest
	GetId() *int64
}

type DeleteDcdnWafGroupRequest struct {
	// The ID of the custom WAF rule group.
	//
	// example:
	//
	// 30000135
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDcdnWafGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDcdnWafGroupRequest) SetId(v int64) *DeleteDcdnWafGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteDcdnWafGroupRequest) Validate() error {
	return dara.Validate(s)
}
