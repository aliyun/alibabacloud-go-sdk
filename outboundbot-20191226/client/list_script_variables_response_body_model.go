// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptVariablesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScriptVariablesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptVariablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScriptVariablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScriptVariablesResponseBody
	GetSuccess() *bool
	SetVariables(v []*ListScriptVariablesResponseBodyVariables) *ListScriptVariablesResponseBody
	GetVariables() []*ListScriptVariablesResponseBodyVariables
}

type ListScriptVariablesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Variables []*ListScriptVariablesResponseBodyVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListScriptVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptVariablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptVariablesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptVariablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptVariablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptVariablesResponseBody) GetVariables() []*ListScriptVariablesResponseBodyVariables {
	return s.Variables
}

func (s *ListScriptVariablesResponseBody) SetCode(v string) *ListScriptVariablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptVariablesResponseBody) SetHttpStatusCode(v int32) *ListScriptVariablesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptVariablesResponseBody) SetMessage(v string) *ListScriptVariablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptVariablesResponseBody) SetRequestId(v string) *ListScriptVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptVariablesResponseBody) SetSuccess(v bool) *ListScriptVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptVariablesResponseBody) SetVariables(v []*ListScriptVariablesResponseBodyVariables) *ListScriptVariablesResponseBody {
	s.Variables = v
	return s
}

func (s *ListScriptVariablesResponseBody) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptVariablesResponseBodyVariables struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListScriptVariablesResponseBodyVariables) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVariablesResponseBodyVariables) GoString() string {
	return s.String()
}

func (s *ListScriptVariablesResponseBodyVariables) GetDescription() *string {
	return s.Description
}

func (s *ListScriptVariablesResponseBodyVariables) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListScriptVariablesResponseBodyVariables) GetName() *string {
	return s.Name
}

func (s *ListScriptVariablesResponseBodyVariables) SetDescription(v string) *ListScriptVariablesResponseBodyVariables {
	s.Description = &v
	return s
}

func (s *ListScriptVariablesResponseBodyVariables) SetDisplayName(v string) *ListScriptVariablesResponseBodyVariables {
	s.DisplayName = &v
	return s
}

func (s *ListScriptVariablesResponseBodyVariables) SetName(v string) *ListScriptVariablesResponseBodyVariables {
	s.Name = &v
	return s
}

func (s *ListScriptVariablesResponseBodyVariables) Validate() error {
	return dara.Validate(s)
}
