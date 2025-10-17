// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructMvDetailModel interface {
	dara.Model
	String() string
	GoString() string
	SetBaseTableInfos(v []*OpenStructMvDetailModelBaseTableInfos) *OpenStructMvDetailModel
	GetBaseTableInfos() []*OpenStructMvDetailModelBaseTableInfos
	SetBaseTableNames(v [][]*string) *OpenStructMvDetailModel
	GetBaseTableNames() [][]*string
	SetEnableDelayAlert(v int32) *OpenStructMvDetailModel
	GetEnableDelayAlert() *int32
	SetEnableFailureAlert(v int32) *OpenStructMvDetailModel
	GetEnableFailureAlert() *int32
	SetExplicitHit(v int64) *OpenStructMvDetailModel
	GetExplicitHit() *int64
	SetFirstRefreshTime(v string) *OpenStructMvDetailModel
	GetFirstRefreshTime() *string
	SetImplicitHit(v int64) *OpenStructMvDetailModel
	GetImplicitHit() *int64
	SetIsInactive(v bool) *OpenStructMvDetailModel
	GetIsInactive() *bool
	SetLatencyTolerance(v int32) *OpenStructMvDetailModel
	GetLatencyTolerance() *int32
	SetLocalSize(v int64) *OpenStructMvDetailModel
	GetLocalSize() *int64
	SetQueryRewriteEnabled(v bool) *OpenStructMvDetailModel
	GetQueryRewriteEnabled() *bool
	SetRefreshInterval(v string) *OpenStructMvDetailModel
	GetRefreshInterval() *string
	SetRefreshState(v string) *OpenStructMvDetailModel
	GetRefreshState() *string
	SetRemoteSize(v int64) *OpenStructMvDetailModel
	GetRemoteSize() *int64
	SetResourceGroup(v string) *OpenStructMvDetailModel
	GetResourceGroup() *string
	SetTableEngine(v string) *OpenStructMvDetailModel
	GetTableEngine() *string
	SetUpdatedAt(v string) *OpenStructMvDetailModel
	GetUpdatedAt() *string
}

type OpenStructMvDetailModel struct {
	BaseTableInfos      []*OpenStructMvDetailModelBaseTableInfos `json:"BaseTableInfos,omitempty" xml:"BaseTableInfos,omitempty" type:"Repeated"`
	BaseTableNames      [][]*string                              `json:"BaseTableNames,omitempty" xml:"BaseTableNames,omitempty" type:"Repeated"`
	EnableDelayAlert    *int32                                   `json:"EnableDelayAlert,omitempty" xml:"EnableDelayAlert,omitempty"`
	EnableFailureAlert  *int32                                   `json:"EnableFailureAlert,omitempty" xml:"EnableFailureAlert,omitempty"`
	ExplicitHit         *int64                                   `json:"ExplicitHit,omitempty" xml:"ExplicitHit,omitempty"`
	FirstRefreshTime    *string                                  `json:"FirstRefreshTime,omitempty" xml:"FirstRefreshTime,omitempty"`
	ImplicitHit         *int64                                   `json:"ImplicitHit,omitempty" xml:"ImplicitHit,omitempty"`
	IsInactive          *bool                                    `json:"IsInactive,omitempty" xml:"IsInactive,omitempty"`
	LatencyTolerance    *int32                                   `json:"LatencyTolerance,omitempty" xml:"LatencyTolerance,omitempty"`
	LocalSize           *int64                                   `json:"LocalSize,omitempty" xml:"LocalSize,omitempty"`
	QueryRewriteEnabled *bool                                    `json:"QueryRewriteEnabled,omitempty" xml:"QueryRewriteEnabled,omitempty"`
	RefreshInterval     *string                                  `json:"RefreshInterval,omitempty" xml:"RefreshInterval,omitempty"`
	RefreshState        *string                                  `json:"RefreshState,omitempty" xml:"RefreshState,omitempty"`
	RemoteSize          *int64                                   `json:"RemoteSize,omitempty" xml:"RemoteSize,omitempty"`
	ResourceGroup       *string                                  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	TableEngine         *string                                  `json:"TableEngine,omitempty" xml:"TableEngine,omitempty"`
	UpdatedAt           *string                                  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s OpenStructMvDetailModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMvDetailModel) GoString() string {
	return s.String()
}

func (s *OpenStructMvDetailModel) GetBaseTableInfos() []*OpenStructMvDetailModelBaseTableInfos {
	return s.BaseTableInfos
}

func (s *OpenStructMvDetailModel) GetBaseTableNames() [][]*string {
	return s.BaseTableNames
}

func (s *OpenStructMvDetailModel) GetEnableDelayAlert() *int32 {
	return s.EnableDelayAlert
}

func (s *OpenStructMvDetailModel) GetEnableFailureAlert() *int32 {
	return s.EnableFailureAlert
}

func (s *OpenStructMvDetailModel) GetExplicitHit() *int64 {
	return s.ExplicitHit
}

func (s *OpenStructMvDetailModel) GetFirstRefreshTime() *string {
	return s.FirstRefreshTime
}

func (s *OpenStructMvDetailModel) GetImplicitHit() *int64 {
	return s.ImplicitHit
}

func (s *OpenStructMvDetailModel) GetIsInactive() *bool {
	return s.IsInactive
}

func (s *OpenStructMvDetailModel) GetLatencyTolerance() *int32 {
	return s.LatencyTolerance
}

func (s *OpenStructMvDetailModel) GetLocalSize() *int64 {
	return s.LocalSize
}

func (s *OpenStructMvDetailModel) GetQueryRewriteEnabled() *bool {
	return s.QueryRewriteEnabled
}

func (s *OpenStructMvDetailModel) GetRefreshInterval() *string {
	return s.RefreshInterval
}

func (s *OpenStructMvDetailModel) GetRefreshState() *string {
	return s.RefreshState
}

func (s *OpenStructMvDetailModel) GetRemoteSize() *int64 {
	return s.RemoteSize
}

func (s *OpenStructMvDetailModel) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *OpenStructMvDetailModel) GetTableEngine() *string {
	return s.TableEngine
}

func (s *OpenStructMvDetailModel) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *OpenStructMvDetailModel) SetBaseTableInfos(v []*OpenStructMvDetailModelBaseTableInfos) *OpenStructMvDetailModel {
	s.BaseTableInfos = v
	return s
}

func (s *OpenStructMvDetailModel) SetBaseTableNames(v [][]*string) *OpenStructMvDetailModel {
	s.BaseTableNames = v
	return s
}

func (s *OpenStructMvDetailModel) SetEnableDelayAlert(v int32) *OpenStructMvDetailModel {
	s.EnableDelayAlert = &v
	return s
}

func (s *OpenStructMvDetailModel) SetEnableFailureAlert(v int32) *OpenStructMvDetailModel {
	s.EnableFailureAlert = &v
	return s
}

func (s *OpenStructMvDetailModel) SetExplicitHit(v int64) *OpenStructMvDetailModel {
	s.ExplicitHit = &v
	return s
}

func (s *OpenStructMvDetailModel) SetFirstRefreshTime(v string) *OpenStructMvDetailModel {
	s.FirstRefreshTime = &v
	return s
}

func (s *OpenStructMvDetailModel) SetImplicitHit(v int64) *OpenStructMvDetailModel {
	s.ImplicitHit = &v
	return s
}

func (s *OpenStructMvDetailModel) SetIsInactive(v bool) *OpenStructMvDetailModel {
	s.IsInactive = &v
	return s
}

func (s *OpenStructMvDetailModel) SetLatencyTolerance(v int32) *OpenStructMvDetailModel {
	s.LatencyTolerance = &v
	return s
}

func (s *OpenStructMvDetailModel) SetLocalSize(v int64) *OpenStructMvDetailModel {
	s.LocalSize = &v
	return s
}

func (s *OpenStructMvDetailModel) SetQueryRewriteEnabled(v bool) *OpenStructMvDetailModel {
	s.QueryRewriteEnabled = &v
	return s
}

func (s *OpenStructMvDetailModel) SetRefreshInterval(v string) *OpenStructMvDetailModel {
	s.RefreshInterval = &v
	return s
}

func (s *OpenStructMvDetailModel) SetRefreshState(v string) *OpenStructMvDetailModel {
	s.RefreshState = &v
	return s
}

func (s *OpenStructMvDetailModel) SetRemoteSize(v int64) *OpenStructMvDetailModel {
	s.RemoteSize = &v
	return s
}

func (s *OpenStructMvDetailModel) SetResourceGroup(v string) *OpenStructMvDetailModel {
	s.ResourceGroup = &v
	return s
}

func (s *OpenStructMvDetailModel) SetTableEngine(v string) *OpenStructMvDetailModel {
	s.TableEngine = &v
	return s
}

func (s *OpenStructMvDetailModel) SetUpdatedAt(v string) *OpenStructMvDetailModel {
	s.UpdatedAt = &v
	return s
}

func (s *OpenStructMvDetailModel) Validate() error {
	if s.BaseTableInfos != nil {
		for _, item := range s.BaseTableInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OpenStructMvDetailModelBaseTableInfos struct {
	BaseTableIsMv *bool   `json:"BaseTableIsMv,omitempty" xml:"BaseTableIsMv,omitempty"`
	SchemaName    *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableEngine   *string `json:"TableEngine,omitempty" xml:"TableEngine,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s OpenStructMvDetailModelBaseTableInfos) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMvDetailModelBaseTableInfos) GoString() string {
	return s.String()
}

func (s *OpenStructMvDetailModelBaseTableInfos) GetBaseTableIsMv() *bool {
	return s.BaseTableIsMv
}

func (s *OpenStructMvDetailModelBaseTableInfos) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OpenStructMvDetailModelBaseTableInfos) GetTableEngine() *string {
	return s.TableEngine
}

func (s *OpenStructMvDetailModelBaseTableInfos) GetTableName() *string {
	return s.TableName
}

func (s *OpenStructMvDetailModelBaseTableInfos) SetBaseTableIsMv(v bool) *OpenStructMvDetailModelBaseTableInfos {
	s.BaseTableIsMv = &v
	return s
}

func (s *OpenStructMvDetailModelBaseTableInfos) SetSchemaName(v string) *OpenStructMvDetailModelBaseTableInfos {
	s.SchemaName = &v
	return s
}

func (s *OpenStructMvDetailModelBaseTableInfos) SetTableEngine(v string) *OpenStructMvDetailModelBaseTableInfos {
	s.TableEngine = &v
	return s
}

func (s *OpenStructMvDetailModelBaseTableInfos) SetTableName(v string) *OpenStructMvDetailModelBaseTableInfos {
	s.TableName = &v
	return s
}

func (s *OpenStructMvDetailModelBaseTableInfos) Validate() error {
	return dara.Validate(s)
}
