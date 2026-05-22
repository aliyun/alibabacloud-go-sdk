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
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
