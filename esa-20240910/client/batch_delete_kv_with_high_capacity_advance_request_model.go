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
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
