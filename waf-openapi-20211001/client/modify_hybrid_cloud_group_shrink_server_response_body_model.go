// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupShrinkServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudGroupShrinkServerResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudGroupShrinkServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudGroupShrinkServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupShrinkServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupShrinkServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudGroupShrinkServerResponseBody) SetRequestId(v string) *ModifyHybridCloudGroupShrinkServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerResponseBody) Validate() error {
	return dara.Validate(s)
}
