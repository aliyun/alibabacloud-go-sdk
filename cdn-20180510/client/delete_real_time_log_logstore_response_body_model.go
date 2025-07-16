// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealTimeLogLogstoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRealTimeLogLogstoreResponseBody
	GetRequestId() *string
}

type DeleteRealTimeLogLogstoreResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRealTimeLogLogstoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealTimeLogLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRealTimeLogLogstoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRealTimeLogLogstoreResponseBody) SetRequestId(v string) *DeleteRealTimeLogLogstoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRealTimeLogLogstoreResponseBody) Validate() error {
	return dara.Validate(s)
}
