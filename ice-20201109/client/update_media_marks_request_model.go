// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaMarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaMarksRequest
	GetMediaId() *string
	SetMediaMarks(v string) *UpdateMediaMarksRequest
	GetMediaMarks() *string
}

type UpdateMediaMarksRequest struct {
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53afdf003a******6a16b5feac6402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The marks of the media asset.
	//
	// This parameter is required.
	MediaMarks *string `json:"MediaMarks,omitempty" xml:"MediaMarks,omitempty"`
}

func (s UpdateMediaMarksRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaMarksRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaMarksRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaMarksRequest) GetMediaMarks() *string {
	return s.MediaMarks
}

func (s *UpdateMediaMarksRequest) SetMediaId(v string) *UpdateMediaMarksRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaMarksRequest) SetMediaMarks(v string) *UpdateMediaMarksRequest {
	s.MediaMarks = &v
	return s
}

func (s *UpdateMediaMarksRequest) Validate() error {
	return dara.Validate(s)
}
