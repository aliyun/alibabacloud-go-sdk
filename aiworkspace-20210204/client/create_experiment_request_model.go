// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateExperimentRequest
	GetAccessibility() *string
	SetArtifactUri(v string) *CreateExperimentRequest
	GetArtifactUri() *string
	SetLabels(v []*LabelInfo) *CreateExperimentRequest
	GetLabels() []*LabelInfo
	SetName(v string) *CreateExperimentRequest
	GetName() *string
	SetWorkspaceId(v string) *CreateExperimentRequest
	GetWorkspaceId() *string
}

type CreateExperimentRequest struct {
	// The visibility of the experiment. Valid values:
	//
	// - PUBLIC: Visible to all users.
	//
	// - PRIVATE: Visible only to the creator and the Alibaba Cloud account.
	//
	// Default value: PRIVATE.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The default artifact output path for all tasks associated with the experiment. Currently, only OSS paths are supported.
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou.aliyuncs.com/test
	ArtifactUri *string `json:"ArtifactUri,omitempty" xml:"ArtifactUri,omitempty"`
	// The labels.
	Labels []*LabelInfo `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the experiment. The name must meet the following rules:
	//
	// - The name must start with a lowercase or uppercase letter.
	//
	// - The name can contain uppercase letters, lowercase letters, digits, underscores (_), or hyphens (-).
	//
	// - The name must be 1 to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateExperimentRequest) GetArtifactUri() *string {
	return s.ArtifactUri
}

func (s *CreateExperimentRequest) GetLabels() []*LabelInfo {
	return s.Labels
}

func (s *CreateExperimentRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperimentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateExperimentRequest) SetAccessibility(v string) *CreateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentRequest) SetArtifactUri(v string) *CreateExperimentRequest {
	s.ArtifactUri = &v
	return s
}

func (s *CreateExperimentRequest) SetLabels(v []*LabelInfo) *CreateExperimentRequest {
	s.Labels = v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetWorkspaceId(v string) *CreateExperimentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateExperimentRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
