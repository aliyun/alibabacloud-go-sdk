// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetLiveSnapshotTemplateRequest
	GetTemplateId() *string
}

type GetLiveSnapshotTemplateRequest struct {
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetLiveSnapshotTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveSnapshotTemplateRequest) SetTemplateId(v string) *GetLiveSnapshotTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLiveSnapshotTemplateRequest) Validate() error {
	return dara.Validate(s)
}
