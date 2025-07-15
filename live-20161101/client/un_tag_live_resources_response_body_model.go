// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagLiveResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnTagLiveResourcesResponseBody
	GetRequestId() *string
}

type UnTagLiveResourcesResponseBody struct {
	// example:
	//
	// 97C68796-EB7F-4D41-9D5B-12B909D76508
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagLiveResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnTagLiveResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagLiveResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnTagLiveResourcesResponseBody) SetRequestId(v string) *UnTagLiveResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnTagLiveResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
