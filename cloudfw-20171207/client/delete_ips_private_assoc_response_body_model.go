// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsPrivateAssocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpsPrivateAssocResponseBody
	GetRequestId() *string
}

type DeleteIpsPrivateAssocResponseBody struct {
	// example:
	//
	// 133173B9-8010-5DF5*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpsPrivateAssocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsPrivateAssocResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpsPrivateAssocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpsPrivateAssocResponseBody) SetRequestId(v string) *DeleteIpsPrivateAssocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpsPrivateAssocResponseBody) Validate() error {
	return dara.Validate(s)
}
