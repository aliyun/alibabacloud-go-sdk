// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaMarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *AddMediaMarksRequest
	GetMediaId() *string
	SetMediaMarks(v string) *AddMediaMarksRequest
	GetMediaMarks() *string
}

type AddMediaMarksRequest struct {
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53afdf003a2171ed9c6a16b5feac6402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The mark information. The value must be a JSONArray.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//  {
	//
	//  "MarkStartTime": "12.1",
	//
	//  "MarkEndTime": "12.2",
	//
	//  "MarkTag": "tag-1",
	//
	//  "MarkContent": "content-1",
	//
	//  "MarkStatus": "a"
	//
	//  },{
	//
	//  "MarkStartTime": "12.3",
	//
	//  "MarkEndTime": "12.4",
	//
	//  "MarkTag": "tag-2",
	//
	//  "MarkContent": "content-2",
	//
	//  "MarkStatus": "online"
	//
	//  }
	//
	// ]
	MediaMarks *string `json:"MediaMarks,omitempty" xml:"MediaMarks,omitempty"`
}

func (s AddMediaMarksRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaMarksRequest) GoString() string {
	return s.String()
}

func (s *AddMediaMarksRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *AddMediaMarksRequest) GetMediaMarks() *string {
	return s.MediaMarks
}

func (s *AddMediaMarksRequest) SetMediaId(v string) *AddMediaMarksRequest {
	s.MediaId = &v
	return s
}

func (s *AddMediaMarksRequest) SetMediaMarks(v string) *AddMediaMarksRequest {
	s.MediaMarks = &v
	return s
}

func (s *AddMediaMarksRequest) Validate() error {
	return dara.Validate(s)
}
