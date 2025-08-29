// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureConsistencyCheckJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFeatureConsistencyCheckJobConfigResponseBody
	GetRequestId() *string
}

type UpdateFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *UpdateFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFeatureConsistencyCheckJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
