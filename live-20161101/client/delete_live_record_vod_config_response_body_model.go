// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordVodConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveRecordVodConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveRecordVodConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRecordVodConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordVodConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordVodConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRecordVodConfigResponseBody) SetRequestId(v string) *DeleteLiveRecordVodConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRecordVodConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
