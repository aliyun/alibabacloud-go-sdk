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
	// The dictionary type. Valid values:
	//
	// - IK: IK cold update dictionary.
	//
	// - IK_HOT: IK hot update dictionary.
	//
	// - SYNONYMS: Synonym dictionary.
	//
	// - ALIWS: Alibaba dictionary.
	//
	// This parameter is required.
	//
	// example:
	//
	// IK
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	// The name of the file to filter.
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
