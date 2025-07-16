// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureViewId(v string) *CreateFeatureViewResponseBody
	GetFeatureViewId() *string
	SetRequestId(v string) *CreateFeatureViewResponseBody
	GetRequestId() *string
}

type CreateFeatureViewResponseBody struct {
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewResponseBody) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *CreateFeatureViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFeatureViewResponseBody) SetFeatureViewId(v string) *CreateFeatureViewResponseBody {
	s.FeatureViewId = &v
	return s
}

func (s *CreateFeatureViewResponseBody) SetRequestId(v string) *CreateFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFeatureViewResponseBody) Validate() error {
	return dara.Validate(s)
}
