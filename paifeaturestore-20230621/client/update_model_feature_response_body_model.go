// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateModelFeatureResponseBody
	GetRequestId() *string
}

type UpdateModelFeatureResponseBody struct {
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelFeatureResponseBody) SetRequestId(v string) *UpdateModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
