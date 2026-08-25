// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceRetentionPolicy(v string) *DeleteTaskRequest
	GetResourceRetentionPolicy() *string
}

type DeleteTaskRequest struct {
	// The data retention policy. If this parameter is not specified, the policy is unconfirmed. If the node has resources or the resource status is unknown, the operation returns a confirmation fault. Set this parameter to RETAIN to delete only the node management record and retain the cloud resources.
	ResourceRetentionPolicy *string `json:"resourceRetentionPolicy,omitempty" xml:"resourceRetentionPolicy,omitempty"`
}

func (s DeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskRequest) GetResourceRetentionPolicy() *string {
	return s.ResourceRetentionPolicy
}

func (s *DeleteTaskRequest) SetResourceRetentionPolicy(v string) *DeleteTaskRequest {
	s.ResourceRetentionPolicy = &v
	return s
}

func (s *DeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
