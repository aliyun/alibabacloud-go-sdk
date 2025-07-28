// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateHybridCloudGroupResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateHybridCloudGroupResponseBody
	GetRequestId() *string
}

type CreateHybridCloudGroupResponseBody struct {
	// The ID of the node group.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 48F7C7BA-0932-50EA-89AD-5B0E1***274
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHybridCloudGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudGroupResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateHybridCloudGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridCloudGroupResponseBody) SetData(v int64) *CreateHybridCloudGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateHybridCloudGroupResponseBody) SetRequestId(v string) *CreateHybridCloudGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridCloudGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
