// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBClusterBindingResponseBodyData) *CreateDBClusterBindingResponseBody
	GetData() *CreateDBClusterBindingResponseBodyData
	SetRequestId(v string) *CreateDBClusterBindingResponseBody
	GetRequestId() *string
}

type CreateDBClusterBindingResponseBody struct {
	// The data returned.
	Data *CreateDBClusterBindingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterBindingResponseBody) GetData() *CreateDBClusterBindingResponseBodyData {
	return s.Data
}

func (s *CreateDBClusterBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBClusterBindingResponseBody) SetData(v *CreateDBClusterBindingResponseBodyData) *CreateDBClusterBindingResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBClusterBindingResponseBody) SetRequestId(v string) *CreateDBClusterBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterBindingResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDBClusterBindingResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The numeric ID.
	//
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s CreateDBClusterBindingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterBindingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBClusterBindingResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *CreateDBClusterBindingResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateDBClusterBindingResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *CreateDBClusterBindingResponseBodyData) SetDbClusterId(v string) *CreateDBClusterBindingResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *CreateDBClusterBindingResponseBodyData) SetDbInstanceId(v string) *CreateDBClusterBindingResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDBClusterBindingResponseBodyData) SetDbInstanceName(v string) *CreateDBClusterBindingResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *CreateDBClusterBindingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
