// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *PublishExperimentRequest
	GetFolderId() *string
}

type PublishExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// root
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s PublishExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishExperimentRequest) GoString() string {
	return s.String()
}

func (s *PublishExperimentRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *PublishExperimentRequest) SetFolderId(v string) *PublishExperimentRequest {
	s.FolderId = &v
	return s
}

func (s *PublishExperimentRequest) Validate() error {
	return dara.Validate(s)
}
