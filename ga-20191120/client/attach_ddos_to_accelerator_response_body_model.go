// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDdosToAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosId(v string) *AttachDdosToAcceleratorResponseBody
	GetDdosId() *string
	SetGaId(v string) *AttachDdosToAcceleratorResponseBody
	GetGaId() *string
	SetRequestId(v string) *AttachDdosToAcceleratorResponseBody
	GetRequestId() *string
}

type AttachDdosToAcceleratorResponseBody struct {
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro/Premium instance that is associated with the GA instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The ID of the GA instance that is associated with the Anti-DDoS Pro/Premium instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	GaId *string `json:"GaId,omitempty" xml:"GaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDdosToAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDdosToAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorResponseBody) GetDdosId() *string {
	return s.DdosId
}

func (s *AttachDdosToAcceleratorResponseBody) GetGaId() *string {
	return s.GaId
}

func (s *AttachDdosToAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDdosToAcceleratorResponseBody) SetDdosId(v string) *AttachDdosToAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) SetGaId(v string) *AttachDdosToAcceleratorResponseBody {
	s.GaId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) SetRequestId(v string) *AttachDdosToAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
