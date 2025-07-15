// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordVodConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveRecordVodConfigResponseBody
	GetRequestId() *string
}

type AddLiveRecordVodConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveRecordVodConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordVodConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveRecordVodConfigResponseBody) SetRequestId(v string) *AddLiveRecordVodConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveRecordVodConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
