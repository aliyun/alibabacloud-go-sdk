// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveStreamDelayConfigResponseBody
	GetRequestId() *string
}

type SetLiveStreamDelayConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4C747C97-7ECD-4C61-8A92-67AD806331FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveStreamDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveStreamDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveStreamDelayConfigResponseBody) SetRequestId(v string) *SetLiveStreamDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveStreamDelayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
