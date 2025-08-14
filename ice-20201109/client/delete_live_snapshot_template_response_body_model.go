// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveSnapshotTemplateResponseBody
	GetRequestId() *string
}

type DeleteLiveSnapshotTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveSnapshotTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveSnapshotTemplateResponseBody) SetRequestId(v string) *DeleteLiveSnapshotTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveSnapshotTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
