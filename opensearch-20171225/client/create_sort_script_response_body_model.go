// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSortScriptResponseBody
	GetRequestId() *string
	SetResult(v *CreateSortScriptResponseBodyResult) *CreateSortScriptResponseBody
	GetResult() *CreateSortScriptResponseBodyResult
}

type CreateSortScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	Result *CreateSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSortScriptResponseBody) GetResult() *CreateSortScriptResponseBodyResult {
	return s.Result
}

func (s *CreateSortScriptResponseBody) SetRequestId(v string) *CreateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSortScriptResponseBody) SetResult(v *CreateSortScriptResponseBodyResult) *CreateSortScriptResponseBody {
	s.Result = v
	return s
}

func (s *CreateSortScriptResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSortScriptResponseBodyResult struct {
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The script name.
	//
	// example:
	//
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The script type.
	//
	// example:
	//
	// cava_script
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSortScriptResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponseBodyResult) GetScope() *string {
	return s.Scope
}

func (s *CreateSortScriptResponseBodyResult) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateSortScriptResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateSortScriptResponseBodyResult) SetScope(v string) *CreateSortScriptResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *CreateSortScriptResponseBodyResult) SetScriptName(v string) *CreateSortScriptResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *CreateSortScriptResponseBodyResult) SetType(v string) *CreateSortScriptResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateSortScriptResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
