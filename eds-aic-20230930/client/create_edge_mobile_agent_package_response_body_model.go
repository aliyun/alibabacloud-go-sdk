// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CreateEdgeMobileAgentPackageResponseBody
	GetOrderId() *int64
	SetPackageIds(v []*string) *CreateEdgeMobileAgentPackageResponseBody
	GetPackageIds() []*string
	SetRequestId(v string) *CreateEdgeMobileAgentPackageResponseBody
	GetRequestId() *string
}

type CreateEdgeMobileAgentPackageResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 22326560487****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The list of package IDs.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEdgeMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeMobileAgentPackageResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateEdgeMobileAgentPackageResponseBody) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *CreateEdgeMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeMobileAgentPackageResponseBody) SetOrderId(v int64) *CreateEdgeMobileAgentPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponseBody) SetPackageIds(v []*string) *CreateEdgeMobileAgentPackageResponseBody {
	s.PackageIds = v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponseBody) SetRequestId(v string) *CreateEdgeMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
