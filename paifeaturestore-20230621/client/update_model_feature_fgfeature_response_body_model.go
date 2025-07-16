// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureFGFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateModelFeatureFGFeatureResponseBody
	GetRequestId() *string
}

type UpdateModelFeatureFGFeatureResponseBody struct {
	// example:
	//
	// 7D497816-607C-5B67-97B1-61354B6ACB2B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateModelFeatureFGFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelFeatureFGFeatureResponseBody) SetRequestId(v string) *UpdateModelFeatureFGFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
