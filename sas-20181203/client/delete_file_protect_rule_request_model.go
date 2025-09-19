// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*int64) *DeleteFileProtectRuleRequest
	GetId() []*int64
}

type DeleteFileProtectRuleRequest struct {
	// The IDs of the core file monitoring rules that you want to delete.
	Id []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DeleteFileProtectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectRuleRequest) GetId() []*int64 {
	return s.Id
}

func (s *DeleteFileProtectRuleRequest) SetId(v []*int64) *DeleteFileProtectRuleRequest {
	s.Id = v
	return s
}

func (s *DeleteFileProtectRuleRequest) Validate() error {
	return dara.Validate(s)
}
