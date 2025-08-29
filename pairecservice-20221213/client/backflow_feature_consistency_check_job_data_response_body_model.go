// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackflowFeatureConsistencyCheckJobDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BackflowFeatureConsistencyCheckJobDataResponseBody
	GetRequestId() *string
}

type BackflowFeatureConsistencyCheckJobDataResponseBody struct {
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataResponseBody) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackflowFeatureConsistencyCheckJobDataResponseBody) SetRequestId(v string) *BackflowFeatureConsistencyCheckJobDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponseBody) Validate() error {
	return dara.Validate(s)
}
