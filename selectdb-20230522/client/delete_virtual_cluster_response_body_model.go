// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteVirtualClusterResponseBodyData) *DeleteVirtualClusterResponseBody
	GetData() *DeleteVirtualClusterResponseBodyData
	SetRequestId(v string) *DeleteVirtualClusterResponseBody
	GetRequestId() *string
}

type DeleteVirtualClusterResponseBody struct {
	Data *DeleteVirtualClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualClusterResponseBody) GetData() *DeleteVirtualClusterResponseBodyData {
	return s.Data
}

func (s *DeleteVirtualClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirtualClusterResponseBody) SetData(v *DeleteVirtualClusterResponseBodyData) *DeleteVirtualClusterResponseBody {
	s.Data = v
	return s
}

func (s *DeleteVirtualClusterResponseBody) SetRequestId(v string) *DeleteVirtualClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteVirtualClusterResponseBodyData struct {
	// example:
	//
	// selectdb-cn-wny3li0****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s DeleteVirtualClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteVirtualClusterResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DeleteVirtualClusterResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DeleteVirtualClusterResponseBodyData) SetDbClusterId(v string) *DeleteVirtualClusterResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *DeleteVirtualClusterResponseBodyData) SetDbInstanceId(v string) *DeleteVirtualClusterResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DeleteVirtualClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
