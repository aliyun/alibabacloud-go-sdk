// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVsStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type SetVsStreamsNotifyUrlConfigResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVsStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *SetVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
