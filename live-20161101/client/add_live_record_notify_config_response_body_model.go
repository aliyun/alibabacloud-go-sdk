// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveRecordNotifyConfigResponseBody
	GetRequestId() *string
}

type AddLiveRecordNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveRecordNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveRecordNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveRecordNotifyConfigResponseBody) SetRequestId(v string) *AddLiveRecordNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveRecordNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
