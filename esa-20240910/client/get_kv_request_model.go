// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64(v bool) *GetKvRequest
	GetBase64() *bool
	SetKey(v string) *GetKvRequest
	GetKey() *string
	SetNamespace(v string) *GetKvRequest
	GetNamespace() *string
}

type GetKvRequest struct {
	Base64 *bool `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKvRequest) GoString() string {
	return s.String()
}

func (s *GetKvRequest) GetBase64() *bool {
	return s.Base64
}

func (s *GetKvRequest) GetKey() *string {
	return s.Key
}

func (s *GetKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvRequest) SetBase64(v bool) *GetKvRequest {
	s.Base64 = &v
	return s
}

func (s *GetKvRequest) SetKey(v string) *GetKvRequest {
	s.Key = &v
	return s
}

func (s *GetKvRequest) SetNamespace(v string) *GetKvRequest {
	s.Namespace = &v
	return s
}

func (s *GetKvRequest) Validate() error {
	return dara.Validate(s)
}
