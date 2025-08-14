// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaMarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DeleteMediaMarksResponseBody
	GetMediaId() *string
	SetMediaMarkIds(v string) *DeleteMediaMarksResponseBody
	GetMediaMarkIds() *string
	SetRequestId(v string) *DeleteMediaMarksResponseBody
	GetRequestId() *string
}

type DeleteMediaMarksResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The IDs of the deleted marks separated with commas (,).
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

func (s DeleteMediaMarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaMarksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaMarksResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMediaMarksResponseBody) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *DeleteMediaMarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaMarksResponseBody) SetMediaId(v string) *DeleteMediaMarksResponseBody {
	s.MediaId = &v
	return s
}

func (s *DeleteMediaMarksResponseBody) SetMediaMarkIds(v string) *DeleteMediaMarksResponseBody {
	s.MediaMarkIds = &v
	return s
}

func (s *DeleteMediaMarksResponseBody) SetRequestId(v string) *DeleteMediaMarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaMarksResponseBody) Validate() error {
	return dara.Validate(s)
}
