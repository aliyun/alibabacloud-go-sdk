// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteKvRequest
	GetKey() *string
	SetNamespace(v string) *DeleteKvRequest
	GetNamespace() *string
}

type DeleteKvRequest struct {
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteKvRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvRequest) GoString() string {
	return s.String()
}

func (s *DeleteKvRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteKvRequest) SetKey(v string) *DeleteKvRequest {
	s.Key = &v
	return s
}

func (s *DeleteKvRequest) SetNamespace(v string) *DeleteKvRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteKvRequest) Validate() error {
	return dara.Validate(s)
}
