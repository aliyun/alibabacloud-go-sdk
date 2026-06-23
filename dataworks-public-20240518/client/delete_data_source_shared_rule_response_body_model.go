// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceSharedRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSourceSharedRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataSourceSharedRuleResponseBody
	GetSuccess() *bool
}

type DeleteDataSourceSharedRuleResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 64B-587A-8CED-969E1973887F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the sharing rule was deleted. Valid values:
	//
	// -true: The sharing rule was deleted.
	//
	// -false: The sharing rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceSharedRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetRequestId(v string) *DeleteDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetSuccess(v bool) *DeleteDataSourceSharedRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
