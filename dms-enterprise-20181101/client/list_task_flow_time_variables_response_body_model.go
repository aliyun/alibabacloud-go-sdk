// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowTimeVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTaskFlowTimeVariablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowTimeVariablesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowTimeVariablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowTimeVariablesResponseBody
	GetSuccess() *bool
	SetTimeVariables(v *ListTaskFlowTimeVariablesResponseBodyTimeVariables) *ListTaskFlowTimeVariablesResponseBody
	GetTimeVariables() *ListTaskFlowTimeVariablesResponseBodyTimeVariables
}

type ListTaskFlowTimeVariablesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC12A3BE-149F-5365-AF33-12CC8C963923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time variables for the task flow.
	TimeVariables *ListTaskFlowTimeVariablesResponseBodyTimeVariables `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty" type:"Struct"`
}

func (s ListTaskFlowTimeVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowTimeVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowTimeVariablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowTimeVariablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowTimeVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowTimeVariablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowTimeVariablesResponseBody) GetTimeVariables() *ListTaskFlowTimeVariablesResponseBodyTimeVariables {
	return s.TimeVariables
}

func (s *ListTaskFlowTimeVariablesResponseBody) SetErrorCode(v string) *ListTaskFlowTimeVariablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBody) SetErrorMessage(v string) *ListTaskFlowTimeVariablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBody) SetRequestId(v string) *ListTaskFlowTimeVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBody) SetSuccess(v bool) *ListTaskFlowTimeVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBody) SetTimeVariables(v *ListTaskFlowTimeVariablesResponseBodyTimeVariables) *ListTaskFlowTimeVariablesResponseBody {
	s.TimeVariables = v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBody) Validate() error {
	if s.TimeVariables != nil {
		if err := s.TimeVariables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowTimeVariablesResponseBodyTimeVariables struct {
	TimeVariable []*ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable `json:"TimeVariable,omitempty" xml:"TimeVariable,omitempty" type:"Repeated"`
}

func (s ListTaskFlowTimeVariablesResponseBodyTimeVariables) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowTimeVariablesResponseBodyTimeVariables) GoString() string {
	return s.String()
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariables) GetTimeVariable() []*ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable {
	return s.TimeVariable
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariables) SetTimeVariable(v []*ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) *ListTaskFlowTimeVariablesResponseBodyTimeVariables {
	s.TimeVariable = v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariables) Validate() error {
	if s.TimeVariable != nil {
		for _, item := range s.TimeVariable {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable struct {
	// The name of the time variable.
	//
	// example:
	//
	// time_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The format of the time variable.
	//
	// example:
	//
	// 2018-09-26|+7h
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
}

func (s ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) GoString() string {
	return s.String()
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) GetName() *string {
	return s.Name
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) GetPattern() *string {
	return s.Pattern
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) SetName(v string) *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable {
	s.Name = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) SetPattern(v string) *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable {
	s.Pattern = &v
	return s
}

func (s *ListTaskFlowTimeVariablesResponseBodyTimeVariablesTimeVariable) Validate() error {
	return dara.Validate(s)
}
