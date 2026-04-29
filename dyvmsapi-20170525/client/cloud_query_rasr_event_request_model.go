// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryRasrEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudQueryRasrEventRequest
	GetEnterpriseId() *int64
	SetUniqueId(v string) *CloudQueryRasrEventRequest
	GetUniqueId() *string
}

type CloudQueryRasrEventRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 通话唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s CloudQueryRasrEventRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryRasrEventRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryRasrEventRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryRasrEventRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryRasrEventRequest) SetEnterpriseId(v int64) *CloudQueryRasrEventRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryRasrEventRequest) SetUniqueId(v string) *CloudQueryRasrEventRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryRasrEventRequest) Validate() error {
	return dara.Validate(s)
}
