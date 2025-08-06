// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVodServiceResponseBody
	GetRequestId() *string
}

type ModifyVodServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVodServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVodServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVodServiceResponseBody) SetRequestId(v string) *ModifyVodServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVodServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
