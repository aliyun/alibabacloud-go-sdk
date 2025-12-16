// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterventionDictionaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzerType(v string) *CreateInterventionDictionaryRequest
	GetAnalyzerType() *string
	SetName(v string) *CreateInterventionDictionaryRequest
	GetName() *string
	SetType(v string) *CreateInterventionDictionaryRequest
	GetType() *string
	SetDryRun(v bool) *CreateInterventionDictionaryRequest
	GetDryRun() *bool
}

type CreateInterventionDictionaryRequest struct {
	// The type of the analyzer. Valid values:
	//
	// 	- MODEL: model-based custom analyzer.
	//
	// 	- SYSTEM: system analyzer.
	//
	// 	- USER: custom analyzer.
	//
	// example:
	//
	// SYSTEM
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// ner_dict_ec
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering.
	//
	// 	- synonym: an intervention dictionary for synonym configuration.
	//
	// 	- correction: an intervention dictionary for spelling correction.
	//
	// 	- category_prediction: an intervention dictionary for category prediction.
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER).
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis.
	//
	// 	- suggest_allowlist: a drop-down suggestion whitelist.
	//
	// 	- suggest_denylist: a drop-down suggestion blacklist.
	//
	// 	- hot_allowlist: a top search whitelist.
	//
	// 	- hot_denylist: a top search blacklist.
	//
	// 	- hint_allowlist: a hint whitelist.
	//
	// 	- hint_denylist: a hint blacklist.
	//
	// example:
	//
	// ner
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateInterventionDictionaryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInterventionDictionaryRequest) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryRequest) GetAnalyzerType() *string {
	return s.AnalyzerType
}

func (s *CreateInterventionDictionaryRequest) GetName() *string {
	return s.Name
}

func (s *CreateInterventionDictionaryRequest) GetType() *string {
	return s.Type
}

func (s *CreateInterventionDictionaryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateInterventionDictionaryRequest) SetAnalyzerType(v string) *CreateInterventionDictionaryRequest {
	s.AnalyzerType = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetName(v string) *CreateInterventionDictionaryRequest {
	s.Name = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetType(v string) *CreateInterventionDictionaryRequest {
	s.Type = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetDryRun(v bool) *CreateInterventionDictionaryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) Validate() error {
	return dara.Validate(s)
}
