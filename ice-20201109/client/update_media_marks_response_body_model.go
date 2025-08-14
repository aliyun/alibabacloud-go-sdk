// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaMarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaMarksResponseBody
	GetMediaId() *string
	SetMediaMarkIds(v string) *UpdateMediaMarksResponseBody
	GetMediaMarkIds() *string
	SetRequestId(v string) *UpdateMediaMarksResponseBody
	GetRequestId() *string
}

type UpdateMediaMarksResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 53afdf003a******6a16b5feac6402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The IDs of the successfully modified marks.
	//
	// example:
	//
	// mark-f82d*****4994b0915948ef7e16,mark-3d56d*****4c8fa9ae2a1f9e5d2d60
	MediaMarkIds *string `json:"MediaMarkIds,omitempty" xml:"MediaMarkIds,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaMarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaMarksResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaMarksResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaMarksResponseBody) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *UpdateMediaMarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaMarksResponseBody) SetMediaId(v string) *UpdateMediaMarksResponseBody {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaMarksResponseBody) SetMediaMarkIds(v string) *UpdateMediaMarksResponseBody {
	s.MediaMarkIds = &v
	return s
}

func (s *UpdateMediaMarksResponseBody) SetRequestId(v string) *UpdateMediaMarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaMarksResponseBody) Validate() error {
	return dara.Validate(s)
}
