// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaMarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaMarksRequest
	GetMediaId() *string
	SetMediaMarkIds(v string) *GetMediaMarksRequest
	GetMediaMarkIds() *string
}

type GetMediaMarksRequest struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The mark ID. You can specify multiple mark IDs separated with commas (,).
	//
	// example:
	//
	// mark-f82d*****4994b0915948ef7e16,mark-3d56d*****4c8fa9ae2a1f9e5d2d60
	MediaMarkIds *string `json:"MediaMarkIds,omitempty" xml:"MediaMarkIds,omitempty"`
}

func (s GetMediaMarksRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaMarksRequest) GoString() string {
	return s.String()
}

func (s *GetMediaMarksRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaMarksRequest) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *GetMediaMarksRequest) SetMediaId(v string) *GetMediaMarksRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaMarksRequest) SetMediaMarkIds(v string) *GetMediaMarksRequest {
	s.MediaMarkIds = &v
	return s
}

func (s *GetMediaMarksRequest) Validate() error {
	return dara.Validate(s)
}
