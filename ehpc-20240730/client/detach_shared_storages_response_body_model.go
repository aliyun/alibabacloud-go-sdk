// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSharedStoragesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DetachSharedStoragesResponseBody
	GetClusterId() *string
	SetRequestId(v string) *DetachSharedStoragesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DetachSharedStoragesResponseBody
	GetSuccess() *string
}

type DetachSharedStoragesResponseBody struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachSharedStoragesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DetachSharedStoragesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachSharedStoragesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DetachSharedStoragesResponseBody) SetClusterId(v string) *DetachSharedStoragesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesResponseBody) SetRequestId(v string) *DetachSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachSharedStoragesResponseBody) SetSuccess(v string) *DetachSharedStoragesResponseBody {
	s.Success = &v
	return s
}

func (s *DetachSharedStoragesResponseBody) Validate() error {
	return dara.Validate(s)
}
