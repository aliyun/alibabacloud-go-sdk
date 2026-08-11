// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptProfileTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptProfileTemplatesRequest
	GetInstanceId() *string
	SetNluEngine(v string) *ListScriptProfileTemplatesRequest
	GetNluEngine() *string
}

type ListScriptProfileTemplatesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The NLU engine type.
	//
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
}

func (s ListScriptProfileTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptProfileTemplatesRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListScriptProfileTemplatesRequest) SetInstanceId(v string) *ListScriptProfileTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptProfileTemplatesRequest) SetNluEngine(v string) *ListScriptProfileTemplatesRequest {
	s.NluEngine = &v
	return s
}

func (s *ListScriptProfileTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
