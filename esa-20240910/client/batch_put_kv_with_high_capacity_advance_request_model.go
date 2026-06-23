// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchPutKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchPutKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchPutKvWithHighCapacityAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
