// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildAICoachScriptRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScriptJsonUrl(v string) *BuildAICoachScriptRecordRequest
	GetScriptJsonUrl() *string
}

type BuildAICoachScriptRecordRequest struct {
	// example:
	//
	// https://xxx/data.json
	ScriptJsonUrl *string `json:"scriptJsonUrl,omitempty" xml:"scriptJsonUrl,omitempty"`
}

func (s BuildAICoachScriptRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildAICoachScriptRecordRequest) GoString() string {
	return s.String()
}

func (s *BuildAICoachScriptRecordRequest) GetScriptJsonUrl() *string {
	return s.ScriptJsonUrl
}

func (s *BuildAICoachScriptRecordRequest) SetScriptJsonUrl(v string) *BuildAICoachScriptRecordRequest {
	s.ScriptJsonUrl = &v
	return s
}

func (s *BuildAICoachScriptRecordRequest) Validate() error {
	return dara.Validate(s)
}
