// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudClusterResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudClusterResponseBody) SetRequestId(v string) *ModifyHybridCloudClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
