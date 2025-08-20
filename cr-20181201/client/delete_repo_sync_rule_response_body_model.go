// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoSyncRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRepoSyncRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteRepoSyncRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteRepoSyncRuleResponseBody
	GetRequestId() *string
}

type DeleteRepoSyncRuleResponseBody struct {
	// The return value.
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
	// The ID of the request.
	//
	// example:
	//
	// 72DD4198-1BB9-47A3-BC01-EAD1A6D5E173
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoSyncRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRepoSyncRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteRepoSyncRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepoSyncRuleResponseBody) SetCode(v string) *DeleteRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoSyncRuleResponseBody) SetIsSuccess(v bool) *DeleteRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoSyncRuleResponseBody) SetRequestId(v string) *DeleteRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepoSyncRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
