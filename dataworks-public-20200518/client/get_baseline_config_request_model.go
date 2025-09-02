// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *GetBaselineConfigRequest
	GetBaselineId() *int64
}

type GetBaselineConfigRequest struct {
	// The baseline ID. You can call the [GetNode](https://help.aliyun.com/document_detail/173977.html) operation to query the baseline ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
}

func (s GetBaselineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineConfigRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineConfigRequest) SetBaselineId(v int64) *GetBaselineConfigRequest {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineConfigRequest) Validate() error {
	return dara.Validate(s)
}
