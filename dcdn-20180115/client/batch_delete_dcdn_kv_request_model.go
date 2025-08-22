// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*string) *BatchDeleteDcdnKvRequest
	GetKeys() []*string
	SetNamespace(v string) *BatchDeleteDcdnKvRequest
	GetNamespace() *string
}

type BatchDeleteDcdnKvRequest struct {
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvRequest) GetKeys() []*string {
	return s.Keys
}

func (s *BatchDeleteDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteDcdnKvRequest) SetKeys(v []*string) *BatchDeleteDcdnKvRequest {
	s.Keys = v
	return s
}

func (s *BatchDeleteDcdnKvRequest) SetNamespace(v string) *BatchDeleteDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}
