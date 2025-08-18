// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppLogRiverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *UpdateEdgeContainerAppLogRiverResponseBody
	GetPath() *string
	SetRequestId(v string) *UpdateEdgeContainerAppLogRiverResponseBody
	GetRequestId() *string
	SetStdout(v bool) *UpdateEdgeContainerAppLogRiverResponseBody
	GetStdout() *bool
}

type UpdateEdgeContainerAppLogRiverResponseBody struct {
	// The log path of the container.
	//
	// example:
	//
	// /root/hello.log
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 42DE97FA-45D2-5615-9A31-55D9EC0D7563
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the standard output of the container is collected.
	//
	// example:
	//
	// true
	Stdout *bool `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
}

func (s UpdateEdgeContainerAppLogRiverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppLogRiverResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) GetPath() *string {
	return s.Path
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) GetStdout() *bool {
	return s.Stdout
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) SetPath(v string) *UpdateEdgeContainerAppLogRiverResponseBody {
	s.Path = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) SetRequestId(v string) *UpdateEdgeContainerAppLogRiverResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) SetStdout(v bool) *UpdateEdgeContainerAppLogRiverResponseBody {
	s.Stdout = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverResponseBody) Validate() error {
	return dara.Validate(s)
}
