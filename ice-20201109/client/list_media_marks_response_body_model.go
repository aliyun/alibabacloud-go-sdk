// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaMarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *ListMediaMarksResponseBody
	GetMediaId() *string
	SetMediaMarks(v string) *ListMediaMarksResponseBody
	GetMediaMarks() *string
	SetRequestId(v string) *ListMediaMarksResponseBody
	GetRequestId() *string
}

type ListMediaMarksResponseBody struct {
	// The ID of the media asset.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The marks of the media asset, in the JSONArray format.
	MediaMarks *string `json:"MediaMarks,omitempty" xml:"MediaMarks,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaMarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaMarksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaMarksResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaMarksResponseBody) GetMediaMarks() *string {
	return s.MediaMarks
}

func (s *ListMediaMarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaMarksResponseBody) SetMediaId(v string) *ListMediaMarksResponseBody {
	s.MediaId = &v
	return s
}

func (s *ListMediaMarksResponseBody) SetMediaMarks(v string) *ListMediaMarksResponseBody {
	s.MediaMarks = &v
	return s
}

func (s *ListMediaMarksResponseBody) SetRequestId(v string) *ListMediaMarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaMarksResponseBody) Validate() error {
	return dara.Validate(s)
}
