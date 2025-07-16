// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeatureId(v string) *CreateModelFeatureResponseBody
	GetModelFeatureId() *string
	SetRequestId(v string) *CreateModelFeatureResponseBody
	GetRequestId() *string
}

type CreateModelFeatureResponseBody struct {
	// example:
	//
	// 3
	ModelFeatureId *string `json:"ModelFeatureId,omitempty" xml:"ModelFeatureId,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureResponseBody) GetModelFeatureId() *string {
	return s.ModelFeatureId
}

func (s *CreateModelFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelFeatureResponseBody) SetModelFeatureId(v string) *CreateModelFeatureResponseBody {
	s.ModelFeatureId = &v
	return s
}

func (s *CreateModelFeatureResponseBody) SetRequestId(v string) *CreateModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
