// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchDeleteKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchDeleteKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchDeleteKvWithHighCapacityAdvanceRequest struct {
	// The namespace name specified when you called [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download URL that contains the key-value pairs to be batch deleted, such as an OSS download URL with read permissions.
	//
	// - When you call this operation by using the SDK, the SDK automatically uploads the content to OSS and passes the corresponding URL.
	//
	// - To call this operation directly, upload the JSON payload (in the same format as the BatchDeleteKv body: {"Namespace":"...","Keys":[...]}) to an OSS bucket and generate a signed HTTPS GET URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-region.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
