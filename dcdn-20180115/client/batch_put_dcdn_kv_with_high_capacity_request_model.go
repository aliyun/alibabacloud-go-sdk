// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutDcdnKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchPutDcdnKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchPutDcdnKvWithHighCapacityRequest struct {
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

func (s BatchPutDcdnKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutDcdnKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchPutDcdnKvWithHighCapacityRequest) SetNamespace(v string) *BatchPutDcdnKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityRequest) SetUrl(v string) *BatchPutDcdnKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
