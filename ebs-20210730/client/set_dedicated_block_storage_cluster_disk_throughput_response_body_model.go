// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDedicatedBlockStorageClusterDiskThroughputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDedicatedBlockStorageClusterDiskThroughputResponseBody
	GetRequestId() *string
}

type SetDedicatedBlockStorageClusterDiskThroughputResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17EE62D8-064E-5404-8B0D-72122478****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponseBody) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) SetRequestId(v string) *SetDedicatedBlockStorageClusterDiskThroughputResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) Validate() error {
	return dara.Validate(s)
}
