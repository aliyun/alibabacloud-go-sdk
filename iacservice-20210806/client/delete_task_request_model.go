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
	// The data retention policy for resources. If this parameter is not specified, it indicates that the deletion is not confirmed. If the node has resources or the resource status is unknown, the API returns a confirmation fault. If you set this parameter to RETAIN, only the node management record is deleted and the cloud resources are retained.
	//
	// example:
	//
	// RETAIN
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
