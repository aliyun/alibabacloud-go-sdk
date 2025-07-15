// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBanLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BanLiveMessageGroupResponseBody
	GetRequestId() *string
}

type BanLiveMessageGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BanLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BanLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BanLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BanLiveMessageGroupResponseBody) SetRequestId(v string) *BanLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BanLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
