// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetStorageACLResponseBody
	GetRequestId() *string
}

type SetStorageACLResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStorageACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetStorageACLResponseBody) GoString() string {
	return s.String()
}

func (s *SetStorageACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetStorageACLResponseBody) SetRequestId(v string) *SetStorageACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetStorageACLResponseBody) Validate() error {
	return dara.Validate(s)
}
