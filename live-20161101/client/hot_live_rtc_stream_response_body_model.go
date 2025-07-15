// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotLiveRtcStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HotLiveRtcStreamResponseBody
	GetRequestId() *string
}

type HotLiveRtcStreamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16BFE188-B193-4C3C-ADC5-79A7E31486EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HotLiveRtcStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotLiveRtcStreamResponseBody) GoString() string {
	return s.String()
}

func (s *HotLiveRtcStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotLiveRtcStreamResponseBody) SetRequestId(v string) *HotLiveRtcStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotLiveRtcStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
