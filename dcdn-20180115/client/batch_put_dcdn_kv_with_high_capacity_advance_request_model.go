// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchPutDcdnKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutDcdnKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchPutDcdnKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchPutDcdnKvWithHighCapacityAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutDcdnKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutDcdnKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchPutDcdnKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchPutDcdnKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchPutDcdnKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
