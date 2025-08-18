// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppLogRiverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *GetEdgeContainerAppLogRiverResponseBody
	GetPath() *string
	SetRequestId(v string) *GetEdgeContainerAppLogRiverResponseBody
	GetRequestId() *string
	SetStdout(v bool) *GetEdgeContainerAppLogRiverResponseBody
	GetStdout() *bool
}

type GetEdgeContainerAppLogRiverResponseBody struct {
	// The log path of the container. It must be an absolute path that starts with a forward slash (/). You can use asterisks (\\*) and question marks (?) as wildcards.
	//
	// example:
	//
	// /root/hello.log
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the standard output of the container is collected.
	Stdout *bool `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
}

func (s GetEdgeContainerAppLogRiverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppLogRiverResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppLogRiverResponseBody) GetPath() *string {
	return s.Path
}

func (s *GetEdgeContainerAppLogRiverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppLogRiverResponseBody) GetStdout() *bool {
	return s.Stdout
}

func (s *GetEdgeContainerAppLogRiverResponseBody) SetPath(v string) *GetEdgeContainerAppLogRiverResponseBody {
	s.Path = &v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponseBody) SetRequestId(v string) *GetEdgeContainerAppLogRiverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponseBody) SetStdout(v bool) *GetEdgeContainerAppLogRiverResponseBody {
	s.Stdout = &v
	return s
}

func (s *GetEdgeContainerAppLogRiverResponseBody) Validate() error {
	return dara.Validate(s)
}
