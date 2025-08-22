// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *GetDcdnKvDetailRequest
	GetKey() *string
	SetNamespace(v string) *GetDcdnKvDetailRequest
	GetNamespace() *string
}

type GetDcdnKvDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetDcdnKvDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDcdnKvDetailRequest) GetKey() *string {
	return s.Key
}

func (s *GetDcdnKvDetailRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDcdnKvDetailRequest) SetKey(v string) *GetDcdnKvDetailRequest {
	s.Key = &v
	return s
}

func (s *GetDcdnKvDetailRequest) SetNamespace(v string) *GetDcdnKvDetailRequest {
	s.Namespace = &v
	return s
}

func (s *GetDcdnKvDetailRequest) Validate() error {
	return dara.Validate(s)
}
