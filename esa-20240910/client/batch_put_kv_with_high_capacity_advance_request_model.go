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
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
