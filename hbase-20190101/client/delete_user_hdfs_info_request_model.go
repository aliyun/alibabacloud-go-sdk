// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserHdfsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUserHdfsInfoRequest
	GetClusterId() *string
	SetNameService(v string) *DeleteUserHdfsInfoRequest
	GetNameService() *string
}

type DeleteUserHdfsInfoRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The value of dfs.nameservices in addUserHdfsInfo. This value is returned when you call the [QueryXpackRelateDB](https://help.aliyun.com/document_detail/144509.html) operation with relateDB set to hdfs.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdfs_test
	NameService *string `json:"NameService,omitempty" xml:"NameService,omitempty"`
}

func (s DeleteUserHdfsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUserHdfsInfoRequest) GetNameService() *string {
	return s.NameService
}

func (s *DeleteUserHdfsInfoRequest) SetClusterId(v string) *DeleteUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetNameService(v string) *DeleteUserHdfsInfoRequest {
	s.NameService = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) Validate() error {
	return dara.Validate(s)
}
