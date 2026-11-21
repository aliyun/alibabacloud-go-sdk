// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElements(v []interface{}) *UpsertUmodelDataRequest
	GetElements() []interface{}
	SetDryRun(v bool) *UpsertUmodelDataRequest
	GetDryRun() *bool
	SetMethod(v string) *UpsertUmodelDataRequest
	GetMethod() *string
}

type UpsertUmodelDataRequest struct {
	// The content of the elements.
	Elements []interface{} `json:"elements,omitempty" xml:"elements,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run for this request. Valid values:
	//
	// - true: sends a check request without performing actual changes.
	//
	// - false: sends a normal request.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The method.
	//
	// example:
	//
	// Upsert
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
}

func (s UpsertUmodelDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelDataRequest) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataRequest) GetElements() []interface{} {
	return s.Elements
}

func (s *UpsertUmodelDataRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpsertUmodelDataRequest) GetMethod() *string {
	return s.Method
}

func (s *UpsertUmodelDataRequest) SetElements(v []interface{}) *UpsertUmodelDataRequest {
	s.Elements = v
	return s
}

func (s *UpsertUmodelDataRequest) SetDryRun(v bool) *UpsertUmodelDataRequest {
	s.DryRun = &v
	return s
}

func (s *UpsertUmodelDataRequest) SetMethod(v string) *UpsertUmodelDataRequest {
	s.Method = &v
	return s
}

func (s *UpsertUmodelDataRequest) Validate() error {
	return dara.Validate(s)
}
