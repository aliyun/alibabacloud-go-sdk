// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScriptRecordId(v string) *GetAICoachScriptRequest
	GetScriptRecordId() *string
}

type GetAICoachScriptRequest struct {
	// example:
	//
	// 1
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
}

func (s GetAICoachScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *GetAICoachScriptRequest) SetScriptRecordId(v string) *GetAICoachScriptRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *GetAICoachScriptRequest) Validate() error {
	return dara.Validate(s)
}
