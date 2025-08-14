// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaMarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaMarksResponseBody
	GetMediaId() *string
	SetMediaMarks(v string) *GetMediaMarksResponseBody
	GetMediaMarks() *string
	SetRequestId(v string) *GetMediaMarksResponseBody
	GetRequestId() *string
}

type GetMediaMarksResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The queried marks.
	//
	// 	- The value is in the JSONArray format.
	MediaMarks *string `json:"MediaMarks,omitempty" xml:"MediaMarks,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaMarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaMarksResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaMarksResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaMarksResponseBody) GetMediaMarks() *string {
	return s.MediaMarks
}

func (s *GetMediaMarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaMarksResponseBody) SetMediaId(v string) *GetMediaMarksResponseBody {
	s.MediaId = &v
	return s
}

func (s *GetMediaMarksResponseBody) SetMediaMarks(v string) *GetMediaMarksResponseBody {
	s.MediaMarks = &v
	return s
}

func (s *GetMediaMarksResponseBody) SetRequestId(v string) *GetMediaMarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaMarksResponseBody) Validate() error {
	return dara.Validate(s)
}
