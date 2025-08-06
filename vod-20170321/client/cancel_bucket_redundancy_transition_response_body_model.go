// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBucketRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelBucketRedundancyTransitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelBucketRedundancyTransitionResponseBody
	GetSuccess() *bool
}

type CancelBucketRedundancyTransitionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelBucketRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelBucketRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBucketRedundancyTransitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelBucketRedundancyTransitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelBucketRedundancyTransitionResponseBody) SetRequestId(v string) *CancelBucketRedundancyTransitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBucketRedundancyTransitionResponseBody) SetSuccess(v bool) *CancelBucketRedundancyTransitionResponseBody {
	s.Success = &v
	return s
}

func (s *CancelBucketRedundancyTransitionResponseBody) Validate() error {
	return dara.Validate(s)
}
