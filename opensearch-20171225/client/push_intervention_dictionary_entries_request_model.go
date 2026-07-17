// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushInterventionDictionaryEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []map[string]interface{}) *PushInterventionDictionaryEntriesRequest
	GetBody() []map[string]interface{}
	SetDryRun(v bool) *PushInterventionDictionaryEntriesRequest
	GetDryRun() *bool
}

type PushInterventionDictionaryEntriesRequest struct {
	// The request body.
	Body []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// Specifies whether to validate the request parameters without creating the attribution configuration. The default value is false.
	//
	// Valid values:
	//
	// - **true**: Validates the request parameters only.
	//
	// - **false**: Validates the request parameters and creates the attribution configuration.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s PushInterventionDictionaryEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s PushInterventionDictionaryEntriesRequest) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesRequest) GetBody() []map[string]interface{} {
	return s.Body
}

func (s *PushInterventionDictionaryEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PushInterventionDictionaryEntriesRequest) SetBody(v []map[string]interface{}) *PushInterventionDictionaryEntriesRequest {
	s.Body = v
	return s
}

func (s *PushInterventionDictionaryEntriesRequest) SetDryRun(v bool) *PushInterventionDictionaryEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *PushInterventionDictionaryEntriesRequest) Validate() error {
	return dara.Validate(s)
}
