// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTaskInstanceDependenciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTaskInstanceDependenciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveTaskInstanceDependenciesResponseBody
	GetSuccess() *bool
}

type RemoveTaskInstanceDependenciesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveTaskInstanceDependenciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTaskInstanceDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTaskInstanceDependenciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTaskInstanceDependenciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveTaskInstanceDependenciesResponseBody) SetRequestId(v string) *RemoveTaskInstanceDependenciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesResponseBody) SetSuccess(v bool) *RemoveTaskInstanceDependenciesResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesResponseBody) Validate() error {
	return dara.Validate(s)
}
