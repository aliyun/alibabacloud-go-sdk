// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAIAgentVoiceprintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearAIAgentVoiceprintResponseBody
	GetRequestId() *string
}

type ClearAIAgentVoiceprintResponseBody struct {
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearAIAgentVoiceprintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearAIAgentVoiceprintResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAIAgentVoiceprintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearAIAgentVoiceprintResponseBody) SetRequestId(v string) *ClearAIAgentVoiceprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearAIAgentVoiceprintResponseBody) Validate() error {
	return dara.Validate(s)
}
