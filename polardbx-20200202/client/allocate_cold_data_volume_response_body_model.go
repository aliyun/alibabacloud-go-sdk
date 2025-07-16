// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateColdDataVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateColdDataVolumeResponseBody
	GetRequestId() *string
}

type AllocateColdDataVolumeResponseBody struct {
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateColdDataVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateColdDataVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateColdDataVolumeResponseBody) SetRequestId(v string) *AllocateColdDataVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateColdDataVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}
