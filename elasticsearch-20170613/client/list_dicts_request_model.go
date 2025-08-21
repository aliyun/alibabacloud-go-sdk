// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzerType(v string) *ListDictsRequest
	GetAnalyzerType() *string
	SetName(v string) *ListDictsRequest
	GetName() *string
}

type ListDictsRequest struct {
	// The type of the dictionary. Valid values:
	//
	// 	- IK: IK dictionary after a standard update
	//
	// 	- IK_HOT: IK dictionary after a rolling update
	//
	// 	- SYNONYMS: synonym dictionary
	//
	// 	- ALIWS: Alibaba Cloud dictionary
	//
	// This parameter is required.
	//
	// example:
	//
	// IK
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDictsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDictsRequest) GoString() string {
	return s.String()
}

func (s *ListDictsRequest) GetAnalyzerType() *string {
	return s.AnalyzerType
}

func (s *ListDictsRequest) GetName() *string {
	return s.Name
}

func (s *ListDictsRequest) SetAnalyzerType(v string) *ListDictsRequest {
	s.AnalyzerType = &v
	return s
}

func (s *ListDictsRequest) SetName(v string) *ListDictsRequest {
	s.Name = &v
	return s
}

func (s *ListDictsRequest) Validate() error {
	return dara.Validate(s)
}
