// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureEntityId(v string) *CreateFeatureEntityResponseBody
	GetFeatureEntityId() *string
	SetRequestId(v string) *CreateFeatureEntityResponseBody
	GetRequestId() *string
}

type CreateFeatureEntityResponseBody struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityResponseBody) GetFeatureEntityId() *string {
	return s.FeatureEntityId
}

func (s *CreateFeatureEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFeatureEntityResponseBody) SetFeatureEntityId(v string) *CreateFeatureEntityResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *CreateFeatureEntityResponseBody) SetRequestId(v string) *CreateFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFeatureEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
