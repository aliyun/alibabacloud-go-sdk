// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAICoachScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScriptId(v string) *OfflineAICoachScriptRequest
	GetScriptId() *string
}

type OfflineAICoachScriptRequest struct {
	// example:
	//
	// 1
	ScriptId *string `json:"scriptId,omitempty" xml:"scriptId,omitempty"`
}

func (s OfflineAICoachScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineAICoachScriptRequest) GoString() string {
	return s.String()
}

func (s *OfflineAICoachScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *OfflineAICoachScriptRequest) SetScriptId(v string) *OfflineAICoachScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *OfflineAICoachScriptRequest) Validate() error {
	return dara.Validate(s)
}
