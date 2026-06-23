// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iPutKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PutKvWithHighCapacityAdvanceRequest
	GetKey() *string
	SetNamespace(v string) *PutKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *PutKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type PutKvWithHighCapacityAdvanceRequest struct {
	// The key name to set. The key name can be up to 512 characters in length and cannot contain spaces or backslashes (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name specified when you called the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// A publicly accessible HTTP or HTTPS URL that points to a JSON file containing the key-value pair to set. The server actively downloads the content from this URL.
	//
	// - If you use an SDK, the SDK automatically uploads the file and generates the URL.
	//
	// - In non-SDK scenarios, upload the JSON payload to any publicly accessible HTTP service and specify the URL.
	//
	// The file content pointed to by the URL must be in the following JSON format: {"Namespace":"<namespace>","Key":"<key>","Value":"<value>"}.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-region.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetKey() *string {
	return s.Key
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetKey(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *PutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
