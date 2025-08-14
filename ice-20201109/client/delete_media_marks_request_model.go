// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaMarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DeleteMediaMarksRequest
	GetMediaId() *string
	SetMediaMarkIds(v string) *DeleteMediaMarksRequest
	GetMediaMarkIds() *string
}

type DeleteMediaMarksRequest struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The mark ID. You can specify multiple mark IDs separated with commas (,).
	//
	// If you do not specify MediaMarkIds, all the marks of the media asset are deleted.
	//
	// example:
	//
	// mark-f82d*****4994b0915948ef7e16,mark-3d56d*****4c8fa9ae2a1f9e5d2d60
	MediaMarkIds *string `json:"MediaMarkIds,omitempty" xml:"MediaMarkIds,omitempty"`
}

func (s DeleteMediaMarksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaMarksRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaMarksRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMediaMarksRequest) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *DeleteMediaMarksRequest) SetMediaId(v string) *DeleteMediaMarksRequest {
	s.MediaId = &v
	return s
}

func (s *DeleteMediaMarksRequest) SetMediaMarkIds(v string) *DeleteMediaMarksRequest {
	s.MediaMarkIds = &v
	return s
}

func (s *DeleteMediaMarksRequest) Validate() error {
	return dara.Validate(s)
}
