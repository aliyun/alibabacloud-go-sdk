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
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
