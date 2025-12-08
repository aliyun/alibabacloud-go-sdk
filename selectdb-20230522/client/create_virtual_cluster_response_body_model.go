// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateVirtualClusterResponseBodyData) *CreateVirtualClusterResponseBody
	GetData() *CreateVirtualClusterResponseBodyData
	SetRequestId(v string) *CreateVirtualClusterResponseBody
	GetRequestId() *string
}

type CreateVirtualClusterResponseBody struct {
	Data *CreateVirtualClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5ED62C81-9948-5612-81E1-EA3853752306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirtualClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualClusterResponseBody) GetData() *CreateVirtualClusterResponseBodyData {
	return s.Data
}

func (s *CreateVirtualClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualClusterResponseBody) SetData(v *CreateVirtualClusterResponseBodyData) *CreateVirtualClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateVirtualClusterResponseBody) SetRequestId(v string) *CreateVirtualClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVirtualClusterResponseBodyData struct {
	// example:
	//
	// selectdb-cn-7213cjv2ez-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
}

func (s CreateVirtualClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVirtualClusterResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *CreateVirtualClusterResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateVirtualClusterResponseBodyData) SetDbClusterId(v string) *CreateVirtualClusterResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *CreateVirtualClusterResponseBodyData) SetDbInstanceId(v string) *CreateVirtualClusterResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *CreateVirtualClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
