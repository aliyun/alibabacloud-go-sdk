// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployRevisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDeployRevisionResponseBody
	GetRequestId() *string
	SetRevision(v *CreateDeployRevisionResponseBodyRevision) *CreateDeployRevisionResponseBody
	GetRevision() *CreateDeployRevisionResponseBodyRevision
}

type CreateDeployRevisionResponseBody struct {
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"RevisionId": "rev-0d6c6956faac431c891b", "ApplicationName": "AgentColin3"}
	Revision *CreateDeployRevisionResponseBodyRevision `json:"Revision,omitempty" xml:"Revision,omitempty" type:"Struct"`
}

func (s CreateDeployRevisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeployRevisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeployRevisionResponseBody) GetRevision() *CreateDeployRevisionResponseBodyRevision {
	return s.Revision
}

func (s *CreateDeployRevisionResponseBody) SetRequestId(v string) *CreateDeployRevisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeployRevisionResponseBody) SetRevision(v *CreateDeployRevisionResponseBodyRevision) *CreateDeployRevisionResponseBody {
	s.Revision = v
	return s
}

func (s *CreateDeployRevisionResponseBody) Validate() error {
	if s.Revision != nil {
		if err := s.Revision.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeployRevisionResponseBodyRevision struct {
	// example:
	//
	// AgentColin3
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// Ecs
	DeployResourceType *string `json:"DeployResourceType,omitempty" xml:"DeployResourceType,omitempty"`
	// example:
	//
	// 2026-01-05
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks       *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// example:
	//
	// {"bucketName":"ecs-application-ui-test","objectName":"319137376-pipeline-run-319137376-task-1-cmd-exec.log","regionId":"cn-hangzhou"}
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// rev-0d6c6956faac431c891b
	RevisionId *string `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	// example:
	//
	// Oss
	RevisionType *string `json:"RevisionType,omitempty" xml:"RevisionType,omitempty"`
}

func (s CreateDeployRevisionResponseBodyRevision) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployRevisionResponseBodyRevision) GoString() string {
	return s.String()
}

func (s *CreateDeployRevisionResponseBodyRevision) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateDeployRevisionResponseBodyRevision) GetDeployResourceType() *string {
	return s.DeployResourceType
}

func (s *CreateDeployRevisionResponseBodyRevision) GetDescription() *string {
	return s.Description
}

func (s *CreateDeployRevisionResponseBodyRevision) GetHooks() *string {
	return s.Hooks
}

func (s *CreateDeployRevisionResponseBodyRevision) GetLocation() *string {
	return s.Location
}

func (s *CreateDeployRevisionResponseBodyRevision) GetRevisionId() *string {
	return s.RevisionId
}

func (s *CreateDeployRevisionResponseBodyRevision) GetRevisionType() *string {
	return s.RevisionType
}

func (s *CreateDeployRevisionResponseBodyRevision) SetApplicationName(v string) *CreateDeployRevisionResponseBodyRevision {
	s.ApplicationName = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetDeployResourceType(v string) *CreateDeployRevisionResponseBodyRevision {
	s.DeployResourceType = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetDescription(v string) *CreateDeployRevisionResponseBodyRevision {
	s.Description = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetHooks(v string) *CreateDeployRevisionResponseBodyRevision {
	s.Hooks = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetLocation(v string) *CreateDeployRevisionResponseBodyRevision {
	s.Location = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetRevisionId(v string) *CreateDeployRevisionResponseBodyRevision {
	s.RevisionId = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) SetRevisionType(v string) *CreateDeployRevisionResponseBodyRevision {
	s.RevisionType = &v
	return s
}

func (s *CreateDeployRevisionResponseBodyRevision) Validate() error {
	return dara.Validate(s)
}
