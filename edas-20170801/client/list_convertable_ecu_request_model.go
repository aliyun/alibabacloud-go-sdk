// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConvertableEcuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListConvertableEcuRequest
	GetClusterId() *string
}

type ListConvertableEcuRequest struct {
	// The ID of the cluster. You can call the ListCluster operation to query the cluster ID. For more information, see [ListCluster](https://help.aliyun.com/document_detail/154995.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// b3e3f77b-462e-****-****-bec8727a****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s ListConvertableEcuRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConvertableEcuRequest) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListConvertableEcuRequest) SetClusterId(v string) *ListConvertableEcuRequest {
	s.ClusterId = &v
	return s
}

func (s *ListConvertableEcuRequest) Validate() error {
	return dara.Validate(s)
}
