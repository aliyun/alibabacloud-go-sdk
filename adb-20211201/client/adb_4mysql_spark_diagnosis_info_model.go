// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdb4MysqlSparkDiagnosisInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosisCode(v string) *Adb4MysqlSparkDiagnosisInfo
	GetDiagnosisCode() *string
	SetDiagnosisCodeLabel(v string) *Adb4MysqlSparkDiagnosisInfo
	GetDiagnosisCodeLabel() *string
	SetDiagnosisMsg(v string) *Adb4MysqlSparkDiagnosisInfo
	GetDiagnosisMsg() *string
	SetDiagnosisType(v string) *Adb4MysqlSparkDiagnosisInfo
	GetDiagnosisType() *string
}

type Adb4MysqlSparkDiagnosisInfo struct {
	// The diagnostic code.
	//
	// example:
	//
	// task_long_tail
	DiagnosisCode *string `json:"DiagnosisCode,omitempty" xml:"DiagnosisCode,omitempty"`
	// The diagnostic cause.
	//
	// example:
	//
	// 查询执行耗时较长。
	DiagnosisCodeLabel *string `json:"DiagnosisCodeLabel,omitempty" xml:"DiagnosisCodeLabel,omitempty"`
	// The diagnostic suggestion.
	//
	// example:
	//
	// 检查查询执行计划并优化 SQL。
	DiagnosisMsg *string `json:"DiagnosisMsg,omitempty" xml:"DiagnosisMsg,omitempty"`
	// The diagnostic information type.
	//
	// example:
	//
	// APPLICATION
	DiagnosisType *string `json:"DiagnosisType,omitempty" xml:"DiagnosisType,omitempty"`
}

func (s Adb4MysqlSparkDiagnosisInfo) String() string {
	return dara.Prettify(s)
}

func (s Adb4MysqlSparkDiagnosisInfo) GoString() string {
	return s.String()
}

func (s *Adb4MysqlSparkDiagnosisInfo) GetDiagnosisCode() *string {
	return s.DiagnosisCode
}

func (s *Adb4MysqlSparkDiagnosisInfo) GetDiagnosisCodeLabel() *string {
	return s.DiagnosisCodeLabel
}

func (s *Adb4MysqlSparkDiagnosisInfo) GetDiagnosisMsg() *string {
	return s.DiagnosisMsg
}

func (s *Adb4MysqlSparkDiagnosisInfo) GetDiagnosisType() *string {
	return s.DiagnosisType
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisCode(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisCode = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisCodeLabel(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisCodeLabel = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisMsg(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisMsg = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisType(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisType = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) Validate() error {
	return dara.Validate(s)
}
