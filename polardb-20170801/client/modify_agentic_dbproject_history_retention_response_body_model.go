// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticDBProjectHistoryRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody
	GetDBClusterId() *string
	SetHistoryRetentionSeconds(v int64) *ModifyAgenticDBProjectHistoryRetentionResponseBody
	GetHistoryRetentionSeconds() *int64
	SetProjectId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody
	GetProjectId() *string
	SetRequestId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody
	GetRequestId() *string
	SetTenantId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody
	GetTenantId() *string
}

type ModifyAgenticDBProjectHistoryRetentionResponseBody struct {
	// DBClusterId
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The current effective history retention period of the project. To check the actual recoverable range, query the branch recovery window.
	//
	// example:
	//
	// 604800
	HistoryRetentionSeconds *int64 `json:"HistoryRetentionSeconds,omitempty" xml:"HistoryRetentionSeconds,omitempty"`
	// ProjectId
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// req-123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TenantId
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyAgenticDBProjectHistoryRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticDBProjectHistoryRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) GetHistoryRetentionSeconds() *int64 {
	return s.HistoryRetentionSeconds
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) SetDBClusterId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) SetHistoryRetentionSeconds(v int64) *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	s.HistoryRetentionSeconds = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) SetProjectId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	s.ProjectId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) SetRequestId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) SetTenantId(v string) *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	s.TenantId = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponseBody) Validate() error {
	return dara.Validate(s)
}
