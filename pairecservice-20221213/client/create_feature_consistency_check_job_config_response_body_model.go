// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureConsistencyCheckJobConfigId() *string
	SetRequestId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody
	GetRequestId() *string
}

type CreateFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *CreateFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
