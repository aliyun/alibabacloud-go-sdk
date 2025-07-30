// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBEClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBEClusterAttributeResponseBody
	GetRequestId() *string
}

type ModifyBEClusterAttributeResponseBody struct {
	// example:
	//
	// 58E21E11-90FF-50F8-A615-8DEB193676E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBEClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBEClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBEClusterAttributeResponseBody) SetRequestId(v string) *ModifyBEClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBEClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
