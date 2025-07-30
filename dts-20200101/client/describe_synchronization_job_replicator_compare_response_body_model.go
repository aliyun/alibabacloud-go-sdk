// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobReplicatorCompareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody
	GetSuccess() *string
	SetSynchronizationReplicatorCompareEnable(v bool) *DescribeSynchronizationJobReplicatorCompareResponseBody
	GetSynchronizationReplicatorCompareEnable() *bool
}

type DescribeSynchronizationJobReplicatorCompareResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E6EB407F-C59F-4682-A682-A00FA6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether image matching is enabled. Valid values:
	//
	// 	- **true**: Image matching is enabled.
	//
	// 	- **false**: Image matching is disabled.
	//
	// example:
	//
	// true
	SynchronizationReplicatorCompareEnable *bool `json:"SynchronizationReplicatorCompareEnable,omitempty" xml:"SynchronizationReplicatorCompareEnable,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) GetSynchronizationReplicatorCompareEnable() *bool {
	return s.SynchronizationReplicatorCompareEnable
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetRequestId(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSynchronizationReplicatorCompareEnable(v bool) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.SynchronizationReplicatorCompareEnable = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) Validate() error {
	return dara.Validate(s)
}
