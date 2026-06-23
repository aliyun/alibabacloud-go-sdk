// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchPutKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchPutKvWithHighCapacityRequest struct {
	// The namespace name specified when you called [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// A publicly accessible HTTP(S) URL that points to a JSON file containing the key-value pairs to be batch set. The server actively downloads the content from this URL.
	//
	// - If you use an SDK, the SDK automatically uploads the file and generates the URL.
	//
	// - In non-SDK scenarios, upload the JSON payload to any publicly accessible HTTP service and specify the URL.
	//
	// The file content pointed to by the URL must be in the following JSON format: {"Namespace":"<namespace name>","KvList":[{"Key":"<key>","Value":"<value>"},...]}.If the URL content does not match this format, the API silently returns an empty SuccessKeys array.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-region.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchPutKvWithHighCapacityRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityRequest) SetUrl(v string) *BatchPutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchPutKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
