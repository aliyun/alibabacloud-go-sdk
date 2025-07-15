// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotDetectPornConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveSnapshotDetectPornConfigResponseBody
	GetRequestId() *string
}

type AddLiveSnapshotDetectPornConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveSnapshotDetectPornConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotDetectPornConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotDetectPornConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveSnapshotDetectPornConfigResponseBody) SetRequestId(v string) *AddLiveSnapshotDetectPornConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
