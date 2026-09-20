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
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd5****582s
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// dfs.nameservices specifies the nameservices name of DFS. dfs.ha.namenodes specifies the DFS nodes, such as nn1 and nn2.
	//
	// dfs.namenode.http-address.{dfs.nameservices}.nn1 specifies the port 50070 connection of HDFS nn1.
	//
	// dfs.namenode.http-address.{dfs.nameservices}.nn2 specifies the port 50070 connection of HDFS nn2.
	//
	// dfs.namenode.rpc-address.{dfs.nameservices}.nn1 specifies the port 8020 connection of HDFS nn1.
	//
	// dfs.namenode.rpc-address.{dfs.nameservices}.nn2 specifies the port 8020 connection of HDFS nn2.
	//
	// Port 50070 connections and port 8020 connections exist on each HDFS node. The number of port 50070 and port 8020 connection pairs equals the number of HDFS nodes.
	//
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
