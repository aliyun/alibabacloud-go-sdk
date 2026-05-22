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
	// The cluster name.
	//
	// example:
	//
	// c497b44c2a59f4ae0bd2826edc40a2c6e
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The container name.
	//
	// example:
	//
	// worker0
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the container group.
	//
	// example:
	//
	// 1775b9e0-8463-457e-89e8-fb7b6d125b2e
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// af22f4xxxxxxxxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The information about the shared token.
	//
	// example:
	//
	// af22f4-xxxxx-xxxx-xxxx-xxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
