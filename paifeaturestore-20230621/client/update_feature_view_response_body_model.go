// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFeatureViewResponseBody
	GetRequestId() *string
}

type UpdateFeatureViewResponseBody struct {
	// example:
	//
	// 7D497816-607C-5B67-97B1-61354B6ACB2B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateFeatureViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFeatureViewResponseBody) SetRequestId(v string) *UpdateFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFeatureViewResponseBody) Validate() error {
	return dara.Validate(s)
}
