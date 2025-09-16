// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *ListSchemasRequest
	GetAccessKey() *string
	SetAccessSecret(v string) *ListSchemasRequest
	GetAccessSecret() *string
	SetEndpoint(v string) *ListSchemasRequest
	GetEndpoint() *string
	SetNamespace(v string) *ListSchemasRequest
	GetNamespace() *string
	SetPartition(v string) *ListSchemasRequest
	GetPartition() *string
	SetProject(v string) *ListSchemasRequest
	GetProject() *string
	SetTable(v string) *ListSchemasRequest
	GetTable() *string
	SetType(v string) *ListSchemasRequest
	GetType() *string
}

type ListSchemasRequest struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// The endpoint of the MaxCompute data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The namespace of the SARO data source.
	//
	// example:
	//
	// igraph-cn-tl32wnrhi04
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The shard name.
	//
	// example:
	//
	// dt=20230520
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// start-flask-v3-obcc
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The name of the MaxCompute table that is used as the data source.
	//
	// example:
	//
	// item
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListSchemasRequest) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *ListSchemasRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListSchemasRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSchemasRequest) GetPartition() *string {
	return s.Partition
}

func (s *ListSchemasRequest) GetProject() *string {
	return s.Project
}

func (s *ListSchemasRequest) GetTable() *string {
	return s.Table
}

func (s *ListSchemasRequest) GetType() *string {
	return s.Type
}

func (s *ListSchemasRequest) SetAccessKey(v string) *ListSchemasRequest {
	s.AccessKey = &v
	return s
}

func (s *ListSchemasRequest) SetAccessSecret(v string) *ListSchemasRequest {
	s.AccessSecret = &v
	return s
}

func (s *ListSchemasRequest) SetEndpoint(v string) *ListSchemasRequest {
	s.Endpoint = &v
	return s
}

func (s *ListSchemasRequest) SetNamespace(v string) *ListSchemasRequest {
	s.Namespace = &v
	return s
}

func (s *ListSchemasRequest) SetPartition(v string) *ListSchemasRequest {
	s.Partition = &v
	return s
}

func (s *ListSchemasRequest) SetProject(v string) *ListSchemasRequest {
	s.Project = &v
	return s
}

func (s *ListSchemasRequest) SetTable(v string) *ListSchemasRequest {
	s.Table = &v
	return s
}

func (s *ListSchemasRequest) SetType(v string) *ListSchemasRequest {
	s.Type = &v
	return s
}

func (s *ListSchemasRequest) Validate() error {
	return dara.Validate(s)
}
