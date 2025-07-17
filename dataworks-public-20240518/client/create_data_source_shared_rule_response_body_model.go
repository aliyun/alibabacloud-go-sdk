// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceSharedRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataSourceSharedRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataSourceSharedRuleResponseBody
	GetRequestId() *string
}

type CreateDataSourceSharedRuleResponseBody struct {
	// The sharing rule ID.
	//
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request. It is used to locate logs and troubleshoot problems.
	//
	// example:
	//
	// 46F594E6-84AB-5FA5-8144-6F3D149961E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceSharedRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataSourceSharedRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceSharedRuleResponseBody) SetId(v int64) *CreateDataSourceSharedRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponseBody) SetRequestId(v string) *CreateDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
