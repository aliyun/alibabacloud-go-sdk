// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *GetDcdnKvRequest
	GetKey() *string
	SetNamespace(v string) *GetDcdnKvRequest
	GetNamespace() *string
}

type GetDcdnKvRequest struct {
	// The name of the key that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *GetDcdnKvRequest) GetKey() *string {
	return s.Key
}

func (s *GetDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDcdnKvRequest) SetKey(v string) *GetDcdnKvRequest {
	s.Key = &v
	return s
}

func (s *GetDcdnKvRequest) SetNamespace(v string) *GetDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *GetDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}
