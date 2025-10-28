// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetFunctions(v []*Function) *ListFunctionsOutput
	GetFunctions() []*Function
	SetNextToken(v string) *ListFunctionsOutput
	GetNextToken() *string
}

type ListFunctionsOutput struct {
	Functions []*Function `json:"functions" xml:"functions" type:"Repeated"`
	// example:
	//
	// next_function_name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsOutput) GoString() string {
	return s.String()
}

func (s *ListFunctionsOutput) GetFunctions() []*Function {
	return s.Functions
}

func (s *ListFunctionsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFunctionsOutput) SetFunctions(v []*Function) *ListFunctionsOutput {
	s.Functions = v
	return s
}

func (s *ListFunctionsOutput) SetNextToken(v string) *ListFunctionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListFunctionsOutput) Validate() error {
	if s.Functions != nil {
		for _, item := range s.Functions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
