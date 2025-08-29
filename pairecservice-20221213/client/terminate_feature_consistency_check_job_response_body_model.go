// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateFeatureConsistencyCheckJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateFeatureConsistencyCheckJobResponseBody
	GetRequestId() *string
}

type TerminateFeatureConsistencyCheckJobResponseBody struct {
	// example:
	//
	// A6C01890-54CA-5C49-BC91-AD85A98E4A98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateFeatureConsistencyCheckJobResponseBody) SetRequestId(v string) *TerminateFeatureConsistencyCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponseBody) Validate() error {
	return dara.Validate(s)
}
