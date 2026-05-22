// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PutKvWithHighCapacityRequest
	GetKey() *string
	SetNamespace(v string) *PutKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *PutKvWithHighCapacityRequest
	GetUrl() *string
}

type PutKvWithHighCapacityRequest struct {
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityRequest) GetKey() *string {
	return s.Key
}

func (s *PutKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *PutKvWithHighCapacityRequest) SetKey(v string) *PutKvWithHighCapacityRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetNamespace(v string) *PutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetUrl(v string) *PutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
