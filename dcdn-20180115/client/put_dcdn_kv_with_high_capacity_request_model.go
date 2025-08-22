// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PutDcdnKvWithHighCapacityRequest
	GetKey() *string
	SetNamespace(v string) *PutDcdnKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *PutDcdnKvWithHighCapacityRequest
	GetUrl() *string
}

type PutDcdnKvWithHighCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutDcdnKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *PutDcdnKvWithHighCapacityRequest) GetKey() *string {
	return s.Key
}

func (s *PutDcdnKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutDcdnKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *PutDcdnKvWithHighCapacityRequest) SetKey(v string) *PutDcdnKvWithHighCapacityRequest {
	s.Key = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityRequest) SetNamespace(v string) *PutDcdnKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityRequest) SetUrl(v string) *PutDcdnKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *PutDcdnKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
