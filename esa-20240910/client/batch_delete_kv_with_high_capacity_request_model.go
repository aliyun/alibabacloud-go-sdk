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
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
