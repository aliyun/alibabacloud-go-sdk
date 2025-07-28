// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterBypassStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudClusterBypassStatusResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudClusterBypassStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudClusterBypassStatusResponseBody) SetRequestId(v string) *ModifyHybridCloudClusterBypassStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
