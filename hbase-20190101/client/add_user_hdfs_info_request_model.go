// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserHdfsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddUserHdfsInfoRequest
	GetClientToken() *string
	SetClusterId(v string) *AddUserHdfsInfoRequest
	GetClusterId() *string
	SetExtInfo(v string) *AddUserHdfsInfoRequest
	GetExtInfo() *string
}

type AddUserHdfsInfoRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd5****582s
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"dfs.nameservices":"hdfs_test","dfs.ha.namenodes":"nn1,nn2","dfs.namenode.http-address.hdfs_test.nn1":"TEST-xxx1.com:50070","dfs.namenode.http-address.hdfs_test.nn2":"TEST-xxx2.com:50070","dfs.namenode.rpc-address.hdfs_test.nn1":"TEST-xxx1.com:8020","dfs.namenode.rpc-address.hdfs_test.nn2":"TEST-xxx2.com:8020"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s AddUserHdfsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddUserHdfsInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddUserHdfsInfoRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *AddUserHdfsInfoRequest) SetClientToken(v string) *AddUserHdfsInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetClusterId(v string) *AddUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetExtInfo(v string) *AddUserHdfsInfoRequest {
	s.ExtInfo = &v
	return s
}

func (s *AddUserHdfsInfoRequest) Validate() error {
	return dara.Validate(s)
}
