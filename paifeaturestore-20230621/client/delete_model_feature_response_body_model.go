// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelFeatureResponseBody
	GetRequestId() *string
}

type DeleteModelFeatureResponseBody struct {
	// example:
	//
	// 6B662A64-E4BF-56F8-BF5F-4C63F34EC0A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelFeatureResponseBody) SetRequestId(v string) *DeleteModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
