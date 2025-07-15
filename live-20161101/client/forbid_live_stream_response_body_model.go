// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidLiveStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ForbidLiveStreamResponseBody
	GetRequestId() *string
}

type ForbidLiveStreamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16BFE188-B193-4C3C-ADC5-79A7E31486EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForbidLiveStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForbidLiveStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ForbidLiveStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForbidLiveStreamResponseBody) SetRequestId(v string) *ForbidLiveStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForbidLiveStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
