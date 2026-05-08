// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAICoachScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScriptId(v string) *DeleteAICoachScriptRequest
	GetScriptId() *string
}

type DeleteAICoachScriptRequest struct {
	// example:
	//
	// 1
	ScriptId *string `json:"scriptId,omitempty" xml:"scriptId,omitempty"`
}

func (s DeleteAICoachScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAICoachScriptRequest) GoString() string {
	return s.String()
}

func (s *DeleteAICoachScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteAICoachScriptRequest) SetScriptId(v string) *DeleteAICoachScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteAICoachScriptRequest) Validate() error {
	return dara.Validate(s)
}
