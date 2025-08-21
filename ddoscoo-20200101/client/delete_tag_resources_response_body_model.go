// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTagResourcesResponseBody
	GetRequestId() *string
}

type DeleteTagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagResourcesResponseBody) SetRequestId(v string) *DeleteTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
