// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ScaleClusterNodePoolResponseBody
	GetRequestId() *string
}

type ScaleClusterNodePoolResponseBody struct {
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleClusterNodePoolResponseBody) SetRequestId(v string) *ScaleClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
