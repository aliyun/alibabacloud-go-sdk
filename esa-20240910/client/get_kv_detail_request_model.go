// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *GetKvDetailRequest
	GetKey() *string
	SetNamespace(v string) *GetKvDetailRequest
	GetNamespace() *string
}

type GetKvDetailRequest struct {
	// The name of the key to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name specified when calling [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKvDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKvDetailRequest) GoString() string {
	return s.String()
}

func (s *GetKvDetailRequest) GetKey() *string {
	return s.Key
}

func (s *GetKvDetailRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvDetailRequest) SetKey(v string) *GetKvDetailRequest {
	s.Key = &v
	return s
}

func (s *GetKvDetailRequest) SetNamespace(v string) *GetKvDetailRequest {
	s.Namespace = &v
	return s
}

func (s *GetKvDetailRequest) Validate() error {
	return dara.Validate(s)
}
