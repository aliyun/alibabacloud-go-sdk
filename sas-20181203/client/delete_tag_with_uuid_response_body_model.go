// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagWithUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTagWithUuidResponseBody
	GetRequestId() *string
}

type DeleteTagWithUuidResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagWithUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagWithUuidResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagWithUuidResponseBody) SetRequestId(v string) *DeleteTagWithUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagWithUuidResponseBody) Validate() error {
	return dara.Validate(s)
}
