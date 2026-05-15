// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityRuleListRequest
	GetInstanceId() *string
	SetPageNo(v int32) *GetQualityRuleListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetQualityRuleListRequest
	GetPageSize() *int32
}

type GetQualityRuleListRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetQualityRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityRuleListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityRuleListRequest) SetInstanceId(v string) *GetQualityRuleListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityRuleListRequest) SetPageNo(v int32) *GetQualityRuleListRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityRuleListRequest) SetPageSize(v int32) *GetQualityRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualityRuleListRequest) Validate() error {
	return dara.Validate(s)
}
