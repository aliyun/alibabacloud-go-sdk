// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppInstanceListRequest
	GetAppId() *string
	SetWithNodeInfo(v bool) *DescribeAppInstanceListRequest
	GetWithNodeInfo() *bool
}

type DescribeAppInstanceListRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the ID of the application. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-*****-ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to return the information about the node in which the pod resides.
	//
	// 	- `true`: returns the information about the node in which the pod resides
	//
	// 	- `false`: does not return the information about the node in which the pod resides
	//
	// example:
	//
	// true
	WithNodeInfo *bool `json:"WithNodeInfo,omitempty" xml:"WithNodeInfo,omitempty"`
}

func (s DescribeAppInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppInstanceListRequest) GetWithNodeInfo() *bool {
	return s.WithNodeInfo
}

func (s *DescribeAppInstanceListRequest) SetAppId(v string) *DescribeAppInstanceListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppInstanceListRequest) SetWithNodeInfo(v bool) *DescribeAppInstanceListRequest {
	s.WithNodeInfo = &v
	return s
}

func (s *DescribeAppInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
