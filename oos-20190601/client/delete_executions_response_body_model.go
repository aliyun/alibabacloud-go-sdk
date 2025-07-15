// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExecutionsResponseBody
	GetRequestId() *string
}

type DeleteExecutionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 491DF8C2-34C9-4679-9DB3-4C0F49B129AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExecutionsResponseBody) SetRequestId(v string) *DeleteExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExecutionsResponseBody) Validate() error {
	return dara.Validate(s)
}
