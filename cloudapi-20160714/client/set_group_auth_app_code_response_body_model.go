// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupAuthAppCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetGroupAuthAppCodeResponseBody
	GetRequestId() *string
}

type SetGroupAuthAppCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGroupAuthAppCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetGroupAuthAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupAuthAppCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetGroupAuthAppCodeResponseBody) SetRequestId(v string) *SetGroupAuthAppCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGroupAuthAppCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
