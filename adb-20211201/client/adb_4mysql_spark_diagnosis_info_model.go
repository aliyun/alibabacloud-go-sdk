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
	DiagnosisCode      *string `json:"DiagnosisCode,omitempty" xml:"DiagnosisCode,omitempty"`
	DiagnosisCodeLabel *string `json:"DiagnosisCodeLabel,omitempty" xml:"DiagnosisCodeLabel,omitempty"`
	DiagnosisMsg       *string `json:"DiagnosisMsg,omitempty" xml:"DiagnosisMsg,omitempty"`
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
