// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildEdgeContainerAppStagingEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebuildEdgeContainerAppStagingEnvResponseBody
	GetRequestId() *string
}

type RebuildEdgeContainerAppStagingEnvResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebuildEdgeContainerAppStagingEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildEdgeContainerAppStagingEnvResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildEdgeContainerAppStagingEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildEdgeContainerAppStagingEnvResponseBody) SetRequestId(v string) *RebuildEdgeContainerAppStagingEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildEdgeContainerAppStagingEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
