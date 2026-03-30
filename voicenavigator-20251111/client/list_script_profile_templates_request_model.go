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
}

type ListScriptProfileTemplatesRequest struct {
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *ListScriptProfileTemplatesRequest) SetInstanceId(v string) *ListScriptProfileTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptProfileTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
