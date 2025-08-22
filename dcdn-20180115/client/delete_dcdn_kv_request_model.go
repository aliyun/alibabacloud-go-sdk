// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteDcdnKvRequest
	GetKey() *string
	SetNamespace(v string) *DeleteDcdnKvRequest
	GetNamespace() *string
}

type DeleteDcdnKvRequest struct {
	// The name of the key that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key_1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The namespace that you specify when you call the PutDcdnKvNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteDcdnKvRequest) SetKey(v string) *DeleteDcdnKvRequest {
	s.Key = &v
	return s
}

func (s *DeleteDcdnKvRequest) SetNamespace(v string) *DeleteDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}
