// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoSyncRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoSyncRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoSyncRuleResponseBody
	GetRequestId() *string
	SetSyncRuleId(v string) *CreateRepoSyncRuleResponseBody
	GetSyncRuleId() *string
}

type CreateRepoSyncRuleResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F8A0BA6-7F06-4BAE-B147-10BD6A25****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-gk5p2ns1kzns****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
}

func (s CreateRepoSyncRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoSyncRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoSyncRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoSyncRuleResponseBody) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *CreateRepoSyncRuleResponseBody) SetCode(v string) *CreateRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetIsSuccess(v bool) *CreateRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetRequestId(v string) *CreateRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetSyncRuleId(v string) *CreateRepoSyncRuleResponseBody {
	s.SyncRuleId = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
