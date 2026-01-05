// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PreviewDatasetVersionRequest
	GetId() *string
}

type PreviewDatasetVersionRequest struct {
	// The dataset version ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-datasetVersion:3pXXXb8o0ngr07njhps1
	//
	// :2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PreviewDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *PreviewDatasetVersionRequest) GetId() *string {
	return s.Id
}

func (s *PreviewDatasetVersionRequest) SetId(v string) *PreviewDatasetVersionRequest {
	s.Id = &v
	return s
}

func (s *PreviewDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
