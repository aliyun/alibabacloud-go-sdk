// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveSnapshotTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLiveSnapshotTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateLiveSnapshotTemplateResponseBody
	GetTemplateId() *string
}

type CreateLiveSnapshotTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateLiveSnapshotTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveSnapshotTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveSnapshotTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveSnapshotTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateLiveSnapshotTemplateResponseBody) SetRequestId(v string) *CreateLiveSnapshotTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveSnapshotTemplateResponseBody) SetTemplateId(v string) *CreateLiveSnapshotTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateLiveSnapshotTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
