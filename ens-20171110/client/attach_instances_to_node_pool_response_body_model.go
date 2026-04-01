// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachInstancesToNodePoolResponseBody
	GetRequestId() *string
}

type AttachInstancesToNodePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachInstancesToNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachInstancesToNodePoolResponseBody) SetRequestId(v string) *AttachInstancesToNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstancesToNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
