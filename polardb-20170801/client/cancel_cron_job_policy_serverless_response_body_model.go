// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCronJobPolicyServerlessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCronJobPolicyServerlessResponseBody
	GetRequestId() *string
}

type CancelCronJobPolicyServerlessResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCronJobPolicyServerlessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCronJobPolicyServerlessResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCronJobPolicyServerlessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCronJobPolicyServerlessResponseBody) SetRequestId(v string) *CancelCronJobPolicyServerlessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCronJobPolicyServerlessResponseBody) Validate() error {
	return dara.Validate(s)
}
