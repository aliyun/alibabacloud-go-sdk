// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *GetDcdnKvStatusRequest
	GetKey() *string
	SetNamespace(v string) *GetDcdnKvStatusRequest
	GetNamespace() *string
}

type GetDcdnKvStatusRequest struct {
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

func (s GetDcdnKvStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDcdnKvStatusRequest) GetKey() *string {
	return s.Key
}

func (s *GetDcdnKvStatusRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDcdnKvStatusRequest) SetKey(v string) *GetDcdnKvStatusRequest {
	s.Key = &v
	return s
}

func (s *GetDcdnKvStatusRequest) SetNamespace(v string) *GetDcdnKvStatusRequest {
	s.Namespace = &v
	return s
}

func (s *GetDcdnKvStatusRequest) Validate() error {
	return dara.Validate(s)
}
