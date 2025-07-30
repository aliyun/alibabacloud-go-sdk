// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveGroupResponseBody
	GetRequestId() *string
}

type RemoveGroupResponseBody struct {
	// example:
	//
	// 42FE70D8-4336-471B-8314-CCCFCE41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveGroupResponseBody) SetRequestId(v string) *RemoveGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
