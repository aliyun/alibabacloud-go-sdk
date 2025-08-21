// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEnsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachEnsInstancesResponseBody
	GetRequestId() *string
}

type AttachEnsInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEnsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEnsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEnsInstancesResponseBody) SetRequestId(v string) *AttachEnsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEnsInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
