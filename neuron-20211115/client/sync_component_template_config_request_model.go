// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncComponentTemplateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *SyncComponentTemplateConfigRequest
	GetId() *int64
}

type SyncComponentTemplateConfigRequest struct {
	// example:
	//
	// 125
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SyncComponentTemplateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncComponentTemplateConfigRequest) GoString() string {
	return s.String()
}

func (s *SyncComponentTemplateConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SyncComponentTemplateConfigRequest) SetId(v int64) *SyncComponentTemplateConfigRequest {
	s.Id = &v
	return s
}

func (s *SyncComponentTemplateConfigRequest) Validate() error {
	return dara.Validate(s)
}
