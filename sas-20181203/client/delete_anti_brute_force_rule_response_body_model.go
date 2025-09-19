// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntiBruteForceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAntiBruteForceRuleResponseBody
	GetRequestId() *string
}

type DeleteAntiBruteForceRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECDF6CE4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAntiBruteForceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntiBruteForceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAntiBruteForceRuleResponseBody) SetRequestId(v string) *DeleteAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAntiBruteForceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
