// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateHybridCloudClusterResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateHybridCloudClusterResponseBody
	GetRequestId() *string
}

type CreateHybridCloudClusterResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-ER12-WE34-23PO-301469*****E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHybridCloudClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateHybridCloudClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridCloudClusterResponseBody) SetData(v int64) *CreateHybridCloudClusterResponseBody {
	s.Data = &v
	return s
}

func (s *CreateHybridCloudClusterResponseBody) SetRequestId(v string) *CreateHybridCloudClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridCloudClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
