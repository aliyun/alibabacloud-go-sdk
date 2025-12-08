// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyVirtualClusterResponseBodyData) *ModifyVirtualClusterResponseBody
	GetData() *ModifyVirtualClusterResponseBodyData
	SetRequestId(v string) *ModifyVirtualClusterResponseBody
	GetRequestId() *string
}

type ModifyVirtualClusterResponseBody struct {
	Data *ModifyVirtualClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVirtualClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVirtualClusterResponseBody) GetData() *ModifyVirtualClusterResponseBodyData {
	return s.Data
}

func (s *ModifyVirtualClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVirtualClusterResponseBody) SetData(v *ModifyVirtualClusterResponseBodyData) *ModifyVirtualClusterResponseBody {
	s.Data = v
	return s
}

func (s *ModifyVirtualClusterResponseBody) SetRequestId(v string) *ModifyVirtualClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVirtualClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyVirtualClusterResponseBodyData struct {
	// example:
	//
	// selectdb-vcg-b****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s ModifyVirtualClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyVirtualClusterResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyVirtualClusterResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyVirtualClusterResponseBodyData) SetDbClusterId(v string) *ModifyVirtualClusterResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *ModifyVirtualClusterResponseBodyData) SetDbInstanceId(v string) *ModifyVirtualClusterResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyVirtualClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
