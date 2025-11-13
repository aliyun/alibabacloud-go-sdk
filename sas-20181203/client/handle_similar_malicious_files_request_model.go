// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarMaliciousFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *HandleSimilarMaliciousFilesRequest
	GetEventId() *int64
	SetLang(v string) *HandleSimilarMaliciousFilesRequest
	GetLang() *string
	SetOperation(v string) *HandleSimilarMaliciousFilesRequest
	GetOperation() *string
	SetScanRange(v string) *HandleSimilarMaliciousFilesRequest
	GetScanRange() *string
	SetScenario(v string) *HandleSimilarMaliciousFilesRequest
	GetScenario() *string
}

type HandleSimilarMaliciousFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ignore
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// agentless
	ScanRange *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// example:
	//
	// same_file_md5
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s HandleSimilarMaliciousFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarMaliciousFilesRequest) GoString() string {
	return s.String()
}

func (s *HandleSimilarMaliciousFilesRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *HandleSimilarMaliciousFilesRequest) GetLang() *string {
	return s.Lang
}

func (s *HandleSimilarMaliciousFilesRequest) GetOperation() *string {
	return s.Operation
}

func (s *HandleSimilarMaliciousFilesRequest) GetScanRange() *string {
	return s.ScanRange
}

func (s *HandleSimilarMaliciousFilesRequest) GetScenario() *string {
	return s.Scenario
}

func (s *HandleSimilarMaliciousFilesRequest) SetEventId(v int64) *HandleSimilarMaliciousFilesRequest {
	s.EventId = &v
	return s
}

func (s *HandleSimilarMaliciousFilesRequest) SetLang(v string) *HandleSimilarMaliciousFilesRequest {
	s.Lang = &v
	return s
}

func (s *HandleSimilarMaliciousFilesRequest) SetOperation(v string) *HandleSimilarMaliciousFilesRequest {
	s.Operation = &v
	return s
}

func (s *HandleSimilarMaliciousFilesRequest) SetScanRange(v string) *HandleSimilarMaliciousFilesRequest {
	s.ScanRange = &v
	return s
}

func (s *HandleSimilarMaliciousFilesRequest) SetScenario(v string) *HandleSimilarMaliciousFilesRequest {
	s.Scenario = &v
	return s
}

func (s *HandleSimilarMaliciousFilesRequest) Validate() error {
	return dara.Validate(s)
}
