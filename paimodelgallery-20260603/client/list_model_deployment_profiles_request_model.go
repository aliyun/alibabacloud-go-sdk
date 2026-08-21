// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelDeploymentProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizKey(v string) *ListModelDeploymentProfilesRequest
	GetBizKey() *string
	SetModelVersion(v string) *ListModelDeploymentProfilesRequest
	GetModelVersion() *string
}

type ListModelDeploymentProfilesRequest struct {
	// example:
	//
	// cmu-biz
	BizKey *string `json:"BizKey,omitempty" xml:"BizKey,omitempty"`
	// example:
	//
	// 1.0.0
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
}

func (s ListModelDeploymentProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesRequest) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesRequest) GetBizKey() *string {
	return s.BizKey
}

func (s *ListModelDeploymentProfilesRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *ListModelDeploymentProfilesRequest) SetBizKey(v string) *ListModelDeploymentProfilesRequest {
	s.BizKey = &v
	return s
}

func (s *ListModelDeploymentProfilesRequest) SetModelVersion(v string) *ListModelDeploymentProfilesRequest {
	s.ModelVersion = &v
	return s
}

func (s *ListModelDeploymentProfilesRequest) Validate() error {
	return dara.Validate(s)
}
