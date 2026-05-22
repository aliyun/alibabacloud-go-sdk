// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetEdgeContainerTerminalResponseBody
	GetCluster() *string
	SetContainer(v string) *GetEdgeContainerTerminalResponseBody
	GetContainer() *string
	SetNamespace(v string) *GetEdgeContainerTerminalResponseBody
	GetNamespace() *string
	SetPod(v string) *GetEdgeContainerTerminalResponseBody
	GetPod() *string
	SetRequestId(v string) *GetEdgeContainerTerminalResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetEdgeContainerTerminalResponseBody
	GetSessionId() *string
	SetToken(v string) *GetEdgeContainerTerminalResponseBody
	GetToken() *string
}

type GetEdgeContainerTerminalResponseBody struct {
	Cluster   *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Pod       *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetEdgeContainerTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerTerminalResponseBody) GetCluster() *string {
	return s.Cluster
}

func (s *GetEdgeContainerTerminalResponseBody) GetContainer() *string {
	return s.Container
}

func (s *GetEdgeContainerTerminalResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *GetEdgeContainerTerminalResponseBody) GetPod() *string {
	return s.Pod
}

func (s *GetEdgeContainerTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerTerminalResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetEdgeContainerTerminalResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetEdgeContainerTerminalResponseBody) SetCluster(v string) *GetEdgeContainerTerminalResponseBody {
	s.Cluster = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetContainer(v string) *GetEdgeContainerTerminalResponseBody {
	s.Container = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetNamespace(v string) *GetEdgeContainerTerminalResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetPod(v string) *GetEdgeContainerTerminalResponseBody {
	s.Pod = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetRequestId(v string) *GetEdgeContainerTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetSessionId(v string) *GetEdgeContainerTerminalResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) SetToken(v string) *GetEdgeContainerTerminalResponseBody {
	s.Token = &v
	return s
}

func (s *GetEdgeContainerTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}
