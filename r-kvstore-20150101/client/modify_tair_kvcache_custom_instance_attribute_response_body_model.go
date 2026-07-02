// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairKVCacheCustomInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTairKVCacheCustomInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyTairKVCacheCustomInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTairKVCacheCustomInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairKVCacheCustomInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponseBody) SetRequestId(v string) *ModifyTairKVCacheCustomInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
