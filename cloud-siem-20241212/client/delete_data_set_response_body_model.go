// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSetResponseBody
	GetRequestId() *string
}

type DeleteDataSetResponseBody struct {
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSetResponseBody) SetRequestId(v string) *DeleteDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
