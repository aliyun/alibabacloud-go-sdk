// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteDcdnKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchDeleteDcdnKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchDeleteDcdnKvWithHighCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteDcdnKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteDcdnKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchDeleteDcdnKvWithHighCapacityRequest) SetNamespace(v string) *BatchDeleteDcdnKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityRequest) SetUrl(v string) *BatchDeleteDcdnKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
