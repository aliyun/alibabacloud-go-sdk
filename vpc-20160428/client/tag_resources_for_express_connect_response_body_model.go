// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesForExpressConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourcesForExpressConnectResponseBody
	GetRequestId() *string
}

type TagResourcesForExpressConnectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesForExpressConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesForExpressConnectResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesForExpressConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesForExpressConnectResponseBody) SetRequestId(v string) *TagResourcesForExpressConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesForExpressConnectResponseBody) Validate() error {
	return dara.Validate(s)
}
