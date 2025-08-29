// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneFeatureConsistencyCheckJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureConsistencyCheckId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody
	GetFeatureConsistencyCheckId() *string
	SetRequestId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody
	GetRequestId() *string
}

type CloneFeatureConsistencyCheckJobConfigResponseBody struct {
	// example:
	//
	// 4
	FeatureConsistencyCheckId *string `json:"FeatureConsistencyCheckId,omitempty" xml:"FeatureConsistencyCheckId,omitempty"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) GetFeatureConsistencyCheckId() *string {
	return s.FeatureConsistencyCheckId
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) SetFeatureConsistencyCheckId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody {
	s.FeatureConsistencyCheckId = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) SetRequestId(v string) *CloneFeatureConsistencyCheckJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
