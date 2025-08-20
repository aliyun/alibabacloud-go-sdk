// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogAnalyzeResult interface {
	dara.Model
	String() string
	GoString() string
	SetAppErrorAdvice(v string) *LogAnalyzeResult
	GetAppErrorAdvice() *string
	SetAppErrorCode(v string) *LogAnalyzeResult
	GetAppErrorCode() *string
	SetAppErrorLog(v string) *LogAnalyzeResult
	GetAppErrorLog() *string
}

type LogAnalyzeResult struct {
	AppErrorAdvice *string `json:"AppErrorAdvice,omitempty" xml:"AppErrorAdvice,omitempty"`
	// example:
	//
	// EXCEEDED_QUOTA
	AppErrorCode *string `json:"AppErrorCode,omitempty" xml:"AppErrorCode,omitempty"`
	// example:
	//
	// exception: xxxx
	AppErrorLog *string `json:"AppErrorLog,omitempty" xml:"AppErrorLog,omitempty"`
}

func (s LogAnalyzeResult) String() string {
	return dara.Prettify(s)
}

func (s LogAnalyzeResult) GoString() string {
	return s.String()
}

func (s *LogAnalyzeResult) GetAppErrorAdvice() *string {
	return s.AppErrorAdvice
}

func (s *LogAnalyzeResult) GetAppErrorCode() *string {
	return s.AppErrorCode
}

func (s *LogAnalyzeResult) GetAppErrorLog() *string {
	return s.AppErrorLog
}

func (s *LogAnalyzeResult) SetAppErrorAdvice(v string) *LogAnalyzeResult {
	s.AppErrorAdvice = &v
	return s
}

func (s *LogAnalyzeResult) SetAppErrorCode(v string) *LogAnalyzeResult {
	s.AppErrorCode = &v
	return s
}

func (s *LogAnalyzeResult) SetAppErrorLog(v string) *LogAnalyzeResult {
	s.AppErrorLog = &v
	return s
}

func (s *LogAnalyzeResult) Validate() error {
	return dara.Validate(s)
}
