// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteLiveSnapshotTemplateRequest
	GetTemplateId() *string
}

type DeleteLiveSnapshotTemplateRequest struct {
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteLiveSnapshotTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteLiveSnapshotTemplateRequest) SetTemplateId(v string) *DeleteLiveSnapshotTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteLiveSnapshotTemplateRequest) Validate() error {
	return dara.Validate(s)
}
