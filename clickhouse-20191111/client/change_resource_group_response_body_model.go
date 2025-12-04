// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
}

type ChangeResourceGroupResponseBody struct {
	// example:
	//
	// C342F3DD-1FF7-55E9-A1A1-098DE07CD1A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
