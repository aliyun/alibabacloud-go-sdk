// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneQuerySceneListByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneName(v string) *DsgSceneQuerySceneListByNameRequest
	GetSceneName() *string
}

type DsgSceneQuerySceneListByNameRequest struct {
	// The name of the data masking scenario. A fuzzy match is performed in the platform based on a keyword to search for the data masking scenario.
	//
	// example:
	//
	// dev_
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s DsgSceneQuerySceneListByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneQuerySceneListByNameRequest) GoString() string {
	return s.String()
}

func (s *DsgSceneQuerySceneListByNameRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *DsgSceneQuerySceneListByNameRequest) SetSceneName(v string) *DsgSceneQuerySceneListByNameRequest {
	s.SceneName = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameRequest) Validate() error {
	return dara.Validate(s)
}
