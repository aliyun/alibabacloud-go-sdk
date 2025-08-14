// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaMarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *ListMediaMarksRequest
	GetMediaId() *string
	SetMediaMarkIds(v string) *ListMediaMarksRequest
	GetMediaMarkIds() *string
}

type ListMediaMarksRequest struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 411bed50018971edb60b0764a0ec6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The mark ID. You can specify multiple IDs separated with commas (,). This parameter is discontinued.
	//
	// example:
	//
	// mark-f82d*****4994b0915948ef7e16,mark-3d56d*****4c8fa9ae2a1f9e5d2d60
	MediaMarkIds *string `json:"MediaMarkIds,omitempty" xml:"MediaMarkIds,omitempty"`
}

func (s ListMediaMarksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaMarksRequest) GoString() string {
	return s.String()
}

func (s *ListMediaMarksRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaMarksRequest) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *ListMediaMarksRequest) SetMediaId(v string) *ListMediaMarksRequest {
	s.MediaId = &v
	return s
}

func (s *ListMediaMarksRequest) SetMediaMarkIds(v string) *ListMediaMarksRequest {
	s.MediaMarkIds = &v
	return s
}

func (s *ListMediaMarksRequest) Validate() error {
	return dara.Validate(s)
}
