// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetObCdrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudGetObCdrRequest
	GetEnterpriseId() *int64
	SetUniqueId(v string) *CloudGetObCdrRequest
	GetUniqueId() *string
}

type CloudGetObCdrRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 电话唯一标识；对应座席外呼通话记录的uniqueId
	//
	// This parameter is required.
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s CloudGetObCdrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetObCdrRequest) GoString() string {
	return s.String()
}

func (s *CloudGetObCdrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetObCdrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudGetObCdrRequest) SetEnterpriseId(v int64) *CloudGetObCdrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetObCdrRequest) SetUniqueId(v string) *CloudGetObCdrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudGetObCdrRequest) Validate() error {
	return dara.Validate(s)
}
