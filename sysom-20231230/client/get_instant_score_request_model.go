// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstantScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetInstantScoreRequest
	GetXDebugId() *string
	SetCluster(v string) *GetInstantScoreRequest
	GetCluster() *string
	SetInstance(v string) *GetInstantScoreRequest
	GetInstance() *string
	SetXSysomInvokeSource(v string) *GetInstantScoreRequest
	GetXSysomInvokeSource() *string
}

type GetInstantScoreRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 2ijff4be-bf24-4070-89ca-c47c879b0g32
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance           *string `json:"instance,omitempty" xml:"instance,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetInstantScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstantScoreRequest) GoString() string {
	return s.String()
}

func (s *GetInstantScoreRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetInstantScoreRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetInstantScoreRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetInstantScoreRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetInstantScoreRequest) SetXDebugId(v string) *GetInstantScoreRequest {
	s.XDebugId = &v
	return s
}

func (s *GetInstantScoreRequest) SetCluster(v string) *GetInstantScoreRequest {
	s.Cluster = &v
	return s
}

func (s *GetInstantScoreRequest) SetInstance(v string) *GetInstantScoreRequest {
	s.Instance = &v
	return s
}

func (s *GetInstantScoreRequest) SetXSysomInvokeSource(v string) *GetInstantScoreRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetInstantScoreRequest) Validate() error {
	return dara.Validate(s)
}
