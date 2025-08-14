// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaMarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *AddMediaMarksResponseBody
	GetMediaId() *string
	SetMediaMarkIds(v string) *AddMediaMarksResponseBody
	GetMediaMarkIds() *string
	SetRequestId(v string) *AddMediaMarksResponseBody
	GetRequestId() *string
}

type AddMediaMarksResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 53afdf003a******6a16b5feac6402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The IDs of the marks that are added.
	//
	// example:
	//
	// mark-f82d*****4994b0915948ef7e16,mark-3d56d*****4c8fa9ae2a1f9e5d2d60
	MediaMarkIds *string `json:"MediaMarkIds,omitempty" xml:"MediaMarkIds,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 771A1414-27BF-53E6-AB73-EFCB*****ACF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMediaMarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaMarksResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaMarksResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *AddMediaMarksResponseBody) GetMediaMarkIds() *string {
	return s.MediaMarkIds
}

func (s *AddMediaMarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaMarksResponseBody) SetMediaId(v string) *AddMediaMarksResponseBody {
	s.MediaId = &v
	return s
}

func (s *AddMediaMarksResponseBody) SetMediaMarkIds(v string) *AddMediaMarksResponseBody {
	s.MediaMarkIds = &v
	return s
}

func (s *AddMediaMarksResponseBody) SetRequestId(v string) *AddMediaMarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaMarksResponseBody) Validate() error {
	return dara.Validate(s)
}
