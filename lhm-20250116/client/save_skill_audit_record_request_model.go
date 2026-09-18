// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSkillAuditRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *SaveSkillAuditRecordRequest
	GetBatchId() *string
	SetCid(v string) *SaveSkillAuditRecordRequest
	GetCid() *string
	SetDryRunStatus(v string) *SaveSkillAuditRecordRequest
	GetDryRunStatus() *string
	SetExtInfo(v string) *SaveSkillAuditRecordRequest
	GetExtInfo() *string
	SetRecordType(v string) *SaveSkillAuditRecordRequest
	GetRecordType() *string
	SetScriptTransformResult(v string) *SaveSkillAuditRecordRequest
	GetScriptTransformResult() *string
	SetScriptTransformStatus(v string) *SaveSkillAuditRecordRequest
	GetScriptTransformStatus() *string
	SetSourceDialect(v string) *SaveSkillAuditRecordRequest
	GetSourceDialect() *string
	SetSourceSqlScript(v string) *SaveSkillAuditRecordRequest
	GetSourceSqlScript() *string
	SetTargetDialect(v string) *SaveSkillAuditRecordRequest
	GetTargetDialect() *string
}

type SaveSkillAuditRecordRequest struct {
	// The batch ID.
	//
	// example:
	//
	// 20001
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 12313
	Cid *string `json:"cid,omitempty" xml:"cid,omitempty"`
	// The dry run status.
	//
	// example:
	//
	// success
	DryRunStatus *string `json:"dryRunStatus,omitempty" xml:"dryRunStatus,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {}
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// The record type.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// The conversion result.
	//
	// example:
	//
	// success
	ScriptTransformResult *string `json:"scriptTransformResult,omitempty" xml:"scriptTransformResult,omitempty"`
	// The conversion status.
	//
	// example:
	//
	// end
	ScriptTransformStatus *string `json:"scriptTransformStatus,omitempty" xml:"scriptTransformStatus,omitempty"`
	// The source dialect.
	//
	// example:
	//
	// hive
	SourceDialect *string `json:"sourceDialect,omitempty" xml:"sourceDialect,omitempty"`
	// The source dialect content.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SourceSqlScript *string `json:"sourceSqlScript,omitempty" xml:"sourceSqlScript,omitempty"`
	// The target dialect.
	//
	// example:
	//
	// bigquery
	TargetDialect *string `json:"targetDialect,omitempty" xml:"targetDialect,omitempty"`
}

func (s SaveSkillAuditRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSkillAuditRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSkillAuditRecordRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *SaveSkillAuditRecordRequest) GetCid() *string {
	return s.Cid
}

func (s *SaveSkillAuditRecordRequest) GetDryRunStatus() *string {
	return s.DryRunStatus
}

func (s *SaveSkillAuditRecordRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *SaveSkillAuditRecordRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *SaveSkillAuditRecordRequest) GetScriptTransformResult() *string {
	return s.ScriptTransformResult
}

func (s *SaveSkillAuditRecordRequest) GetScriptTransformStatus() *string {
	return s.ScriptTransformStatus
}

func (s *SaveSkillAuditRecordRequest) GetSourceDialect() *string {
	return s.SourceDialect
}

func (s *SaveSkillAuditRecordRequest) GetSourceSqlScript() *string {
	return s.SourceSqlScript
}

func (s *SaveSkillAuditRecordRequest) GetTargetDialect() *string {
	return s.TargetDialect
}

func (s *SaveSkillAuditRecordRequest) SetBatchId(v string) *SaveSkillAuditRecordRequest {
	s.BatchId = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetCid(v string) *SaveSkillAuditRecordRequest {
	s.Cid = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetDryRunStatus(v string) *SaveSkillAuditRecordRequest {
	s.DryRunStatus = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetExtInfo(v string) *SaveSkillAuditRecordRequest {
	s.ExtInfo = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetRecordType(v string) *SaveSkillAuditRecordRequest {
	s.RecordType = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetScriptTransformResult(v string) *SaveSkillAuditRecordRequest {
	s.ScriptTransformResult = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetScriptTransformStatus(v string) *SaveSkillAuditRecordRequest {
	s.ScriptTransformStatus = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetSourceDialect(v string) *SaveSkillAuditRecordRequest {
	s.SourceDialect = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetSourceSqlScript(v string) *SaveSkillAuditRecordRequest {
	s.SourceSqlScript = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) SetTargetDialect(v string) *SaveSkillAuditRecordRequest {
	s.TargetDialect = &v
	return s
}

func (s *SaveSkillAuditRecordRequest) Validate() error {
	return dara.Validate(s)
}
