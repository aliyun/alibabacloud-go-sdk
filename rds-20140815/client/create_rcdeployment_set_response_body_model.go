// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *CreateRCDeploymentSetResponseBody
	GetDeploymentSetId() *string
	SetRequestId(v string) *CreateRCDeploymentSetResponseBody
	GetRequestId() *string
}

type CreateRCDeploymentSetResponseBody struct {
	// The deployment set ID.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRCDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCDeploymentSetResponseBody) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateRCDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCDeploymentSetResponseBody) SetDeploymentSetId(v string) *CreateRCDeploymentSetResponseBody {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateRCDeploymentSetResponseBody) SetRequestId(v string) *CreateRCDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCDeploymentSetResponseBody) Validate() error {
	return dara.Validate(s)
}
