// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoSyncRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRepoSyncRuleRequest
	GetInstanceId() *string
	SetSyncRuleId(v string) *DeleteRepoSyncRuleRequest
	GetSyncRuleId() *string
}

type DeleteRepoSyncRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsr-gk5p2ns1kzns****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
}

func (s DeleteRepoSyncRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRepoSyncRuleRequest) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *DeleteRepoSyncRuleRequest) SetInstanceId(v string) *DeleteRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoSyncRuleRequest) SetSyncRuleId(v string) *DeleteRepoSyncRuleRequest {
	s.SyncRuleId = &v
	return s
}

func (s *DeleteRepoSyncRuleRequest) Validate() error {
	return dara.Validate(s)
}
