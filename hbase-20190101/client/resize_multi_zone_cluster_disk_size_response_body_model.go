// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterDiskSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody
	GetRequestId() *string
}

type ResizeMultiZoneClusterDiskSizeResponseBody struct {
	// example:
	//
	// 123412341****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 568339C4-9F71-43D0-994E-E039CD826E56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) Validate() error {
	return dara.Validate(s)
}
