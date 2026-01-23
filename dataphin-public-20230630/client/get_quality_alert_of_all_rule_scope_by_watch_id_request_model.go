// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityAlertOfAllRuleScopeByWatchIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityAlertOfAllRuleScopeByWatchIdRequest
	GetOpTenantId() *int64
	SetWatchId(v int64) *GetQualityAlertOfAllRuleScopeByWatchIdRequest
	GetWatchId() *int64
}

type GetQualityAlertOfAllRuleScopeByWatchIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityAlertOfAllRuleScopeByWatchIdRequest) GoString() string {
	return s.String()
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdRequest) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdRequest) SetOpTenantId(v int64) *GetQualityAlertOfAllRuleScopeByWatchIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdRequest) SetWatchId(v int64) *GetQualityAlertOfAllRuleScopeByWatchIdRequest {
	s.WatchId = &v
	return s
}

func (s *GetQualityAlertOfAllRuleScopeByWatchIdRequest) Validate() error {
	return dara.Validate(s)
}
