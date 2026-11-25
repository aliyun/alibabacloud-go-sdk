// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarmPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWarmPoolResponseBody
	GetRequestId() *string
}

type DeleteWarmPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWarmPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarmPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarmPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWarmPoolResponseBody) SetRequestId(v string) *DeleteWarmPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarmPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
