// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLogstashDeletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelLogstashDeletionResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelLogstashDeletionResponseBody
	GetResult() *bool
}

type CancelLogstashDeletionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 23EBF56B-2DC0-4507-8BE5-B87395DB0FEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the cluster is restored. Valid values:
	//
	// 	- true: The cluster is restored.
	//
	// 	- false: The cluster is not restored.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelLogstashDeletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelLogstashDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelLogstashDeletionResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelLogstashDeletionResponseBody) SetRequestId(v string) *CancelLogstashDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelLogstashDeletionResponseBody) SetResult(v bool) *CancelLogstashDeletionResponseBody {
	s.Result = &v
	return s
}

func (s *CancelLogstashDeletionResponseBody) Validate() error {
	return dara.Validate(s)
}
