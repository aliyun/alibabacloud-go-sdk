// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckJobId(v string) *CreateFeatureConsistencyCheckJobResponseBody
	GetFeatureConsistencyCheckJobId() *string
	SetRequestId(v string) *CreateFeatureConsistencyCheckJobResponseBody
	GetRequestId() *string
}

type CreateFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckJobId *string `json:"FeatureConsistencyCheckJobId,omitempty" xml:"FeatureConsistencyCheckJobId,omitempty"`
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) GetFeatureConsistencyCheckJobId() *string {
	return s.FeatureConsistencyCheckJobId
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) SetFeatureConsistencyCheckJobId(v string) *CreateFeatureConsistencyCheckJobResponseBody {
	s.FeatureConsistencyCheckJobId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *CreateFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponseBody) Validate() error {
	return dara.Validate(s)
}
