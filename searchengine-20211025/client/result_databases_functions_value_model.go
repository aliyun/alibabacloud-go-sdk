// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultDatabasesFunctionsValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ResultDatabasesFunctionsValue
	GetName() *string
	SetSignatures(v string) *ResultDatabasesFunctionsValue
	GetSignatures() *string
}

type ResultDatabasesFunctionsValue struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// gfasdds2****2wfrkv
	Signatures *string `json:"signatures,omitempty" xml:"signatures,omitempty"`
}

func (s ResultDatabasesFunctionsValue) String() string {
	return dara.Prettify(s)
}

func (s ResultDatabasesFunctionsValue) GoString() string {
	return s.String()
}

func (s *ResultDatabasesFunctionsValue) GetName() *string {
	return s.Name
}

func (s *ResultDatabasesFunctionsValue) GetSignatures() *string {
	return s.Signatures
}

func (s *ResultDatabasesFunctionsValue) SetName(v string) *ResultDatabasesFunctionsValue {
	s.Name = &v
	return s
}

func (s *ResultDatabasesFunctionsValue) SetSignatures(v string) *ResultDatabasesFunctionsValue {
	s.Signatures = &v
	return s
}

func (s *ResultDatabasesFunctionsValue) Validate() error {
	return dara.Validate(s)
}
