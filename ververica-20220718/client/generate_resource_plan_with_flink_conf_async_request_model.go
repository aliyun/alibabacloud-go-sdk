// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourcePlanWithFlinkConfAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *GenerateResourcePlanWithFlinkConfAsyncRequest
	GetBody() map[string]interface{}
}

type GenerateResourcePlanWithFlinkConfAsyncRequest struct {
	// The Flink configuration that is used to generate a resource plan.
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncRequest) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GenerateResourcePlanWithFlinkConfAsyncRequest) SetBody(v map[string]interface{}) *GenerateResourcePlanWithFlinkConfAsyncRequest {
	s.Body = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncRequest) Validate() error {
	return dara.Validate(s)
}
