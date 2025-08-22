// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchDeleteDcdnKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchDeleteDcdnKvWithHighCapacityAdvanceRequest struct {
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

func (s BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
