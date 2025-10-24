// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTagGroupResponseBody
	GetSuccess() *bool
}

type DeleteTagGroupResponseBody struct {
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTagGroupResponseBody) SetRequestId(v string) *DeleteTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagGroupResponseBody) SetSuccess(v bool) *DeleteTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
