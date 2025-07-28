// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudGroupResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D***0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudGroupResponseBody) SetRequestId(v string) *ModifyHybridCloudGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
