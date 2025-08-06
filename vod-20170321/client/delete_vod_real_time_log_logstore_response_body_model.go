// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealTimeLogLogstoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodRealTimeLogLogstoreResponseBody
	GetRequestId() *string
}

type DeleteVodRealTimeLogLogstoreResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodRealTimeLogLogstoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealTimeLogLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodRealTimeLogLogstoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodRealTimeLogLogstoreResponseBody) SetRequestId(v string) *DeleteVodRealTimeLogLogstoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreResponseBody) Validate() error {
	return dara.Validate(s)
}
