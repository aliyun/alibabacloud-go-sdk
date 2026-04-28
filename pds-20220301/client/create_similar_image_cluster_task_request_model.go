// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *CreateSimilarImageClusterTaskRequest
	GetDriveId() *string
}

type CreateSimilarImageClusterTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s CreateSimilarImageClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateSimilarImageClusterTaskRequest) SetDriveId(v string) *CreateSimilarImageClusterTaskRequest {
	s.DriveId = &v
	return s
}

func (s *CreateSimilarImageClusterTaskRequest) Validate() error {
	return dara.Validate(s)
}
