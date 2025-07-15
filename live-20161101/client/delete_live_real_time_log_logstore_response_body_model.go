// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealTimeLogLogstoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveRealTimeLogLogstoreResponseBody
	GetRequestId() *string
}

type DeleteLiveRealTimeLogLogstoreResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRealTimeLogLogstoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealTimeLogLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealTimeLogLogstoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRealTimeLogLogstoreResponseBody) SetRequestId(v string) *DeleteLiveRealTimeLogLogstoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreResponseBody) Validate() error {
	return dara.Validate(s)
}
