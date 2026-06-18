// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchDeleteKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchDeleteKvWithHighCapacityRequest struct {
	// The name specified when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download URL that stores the key-value pairs to be batch deleted, such as an OSS download URL with read permission.
	//
	// - When you call this operation by using an SDK, the SDK automatically uploads the submitted content to OSS and passes in the corresponding URL.
	//
	// - To call this operation directly, you must upload the JSON payload (in the same format as the BatchDeleteKv body: {"Namespace":"...","Keys":[...]}) to an OSS bucket and generate a signed HTTPS GET URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-region.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetUrl(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
