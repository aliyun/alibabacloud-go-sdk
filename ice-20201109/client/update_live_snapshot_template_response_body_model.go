// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveSnapshotTemplateResponseBody
	GetRequestId() *string
}

type UpdateLiveSnapshotTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveSnapshotTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveSnapshotTemplateResponseBody) SetRequestId(v string) *UpdateLiveSnapshotTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
