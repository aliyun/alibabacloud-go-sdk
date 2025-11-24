// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentBySelectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentNameList(v [][]byte) *GetDeploymentBySelectorResponseBody
	GetDeploymentNameList() [][]byte
	SetMark(v string) *GetDeploymentBySelectorResponseBody
	GetMark() *string
	SetRequestId(v string) *GetDeploymentBySelectorResponseBody
	GetRequestId() *string
}

type GetDeploymentBySelectorResponseBody struct {
	// The queried workloads.
	DeploymentNameList [][]byte `json:"DeploymentNameList,omitempty" xml:"DeploymentNameList,omitempty" type:"Repeated"`
	// The end-of-data marker.
	//
	// example:
	//
	// eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6NTgyMDUzMzk5MCwic3RhcnQiOiJwbXMtYWRhcHRlci1kZGxsXHUwMDA****
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 946690C2-41D3-55A0-A501-E2FFAB5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeploymentBySelectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentBySelectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorResponseBody) GetDeploymentNameList() [][]byte {
	return s.DeploymentNameList
}

func (s *GetDeploymentBySelectorResponseBody) GetMark() *string {
	return s.Mark
}

func (s *GetDeploymentBySelectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentBySelectorResponseBody) SetDeploymentNameList(v [][]byte) *GetDeploymentBySelectorResponseBody {
	s.DeploymentNameList = v
	return s
}

func (s *GetDeploymentBySelectorResponseBody) SetMark(v string) *GetDeploymentBySelectorResponseBody {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorResponseBody) SetRequestId(v string) *GetDeploymentBySelectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentBySelectorResponseBody) Validate() error {
	return dara.Validate(s)
}
