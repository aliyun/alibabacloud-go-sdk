// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDetectNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveDetectNotifyConfigResponseBody
	GetRequestId() *string
}

type AddLiveDetectNotifyConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveDetectNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDetectNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveDetectNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveDetectNotifyConfigResponseBody) SetRequestId(v string) *AddLiveDetectNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveDetectNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
