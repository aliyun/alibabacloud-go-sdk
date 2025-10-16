// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *RetrieveMemoryRequest
	GetFrom() *int32
	SetQuery(v *RetrieveMemoryRequestQuery) *RetrieveMemoryRequest
	GetQuery() *RetrieveMemoryRequestQuery
	SetStore(v string) *RetrieveMemoryRequest
	GetStore() *string
	SetTo(v int32) *RetrieveMemoryRequest
	GetTo() *int32
	SetTopk(v int32) *RetrieveMemoryRequest
	GetTopk() *int32
}

type RetrieveMemoryRequest struct {
	// example:
	//
	// 1756431013
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	Query *RetrieveMemoryRequestQuery `json:"query,omitempty" xml:"query,omitempty" type:"Struct"`
	// example:
	//
	// memory
	Store *string `json:"store,omitempty" xml:"store,omitempty"`
	// example:
	//
	// 1737856802
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// example:
	//
	// 10
	Topk *int32 `json:"topk,omitempty" xml:"topk,omitempty"`
}

func (s RetrieveMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveMemoryRequest) GoString() string {
	return s.String()
}

func (s *RetrieveMemoryRequest) GetFrom() *int32 {
	return s.From
}

func (s *RetrieveMemoryRequest) GetQuery() *RetrieveMemoryRequestQuery {
	return s.Query
}

func (s *RetrieveMemoryRequest) GetStore() *string {
	return s.Store
}

func (s *RetrieveMemoryRequest) GetTo() *int32 {
	return s.To
}

func (s *RetrieveMemoryRequest) GetTopk() *int32 {
	return s.Topk
}

func (s *RetrieveMemoryRequest) SetFrom(v int32) *RetrieveMemoryRequest {
	s.From = &v
	return s
}

func (s *RetrieveMemoryRequest) SetQuery(v *RetrieveMemoryRequestQuery) *RetrieveMemoryRequest {
	s.Query = v
	return s
}

func (s *RetrieveMemoryRequest) SetStore(v string) *RetrieveMemoryRequest {
	s.Store = &v
	return s
}

func (s *RetrieveMemoryRequest) SetTo(v int32) *RetrieveMemoryRequest {
	s.To = &v
	return s
}

func (s *RetrieveMemoryRequest) SetTopk(v int32) *RetrieveMemoryRequest {
	s.Topk = &v
	return s
}

func (s *RetrieveMemoryRequest) Validate() error {
	if s.Query != nil {
		if err := s.Query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveMemoryRequestQuery struct {
	// example:
	//
	// 10
	Memory   *string            `json:"memory,omitempty" xml:"memory,omitempty"`
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// preference
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uid1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RetrieveMemoryRequestQuery) String() string {
	return dara.Prettify(s)
}

func (s RetrieveMemoryRequestQuery) GoString() string {
	return s.String()
}

func (s *RetrieveMemoryRequestQuery) GetMemory() *string {
	return s.Memory
}

func (s *RetrieveMemoryRequestQuery) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *RetrieveMemoryRequestQuery) GetNamespace() *string {
	return s.Namespace
}

func (s *RetrieveMemoryRequestQuery) GetUserId() *string {
	return s.UserId
}

func (s *RetrieveMemoryRequestQuery) SetMemory(v string) *RetrieveMemoryRequestQuery {
	s.Memory = &v
	return s
}

func (s *RetrieveMemoryRequestQuery) SetMetadata(v map[string]*string) *RetrieveMemoryRequestQuery {
	s.Metadata = v
	return s
}

func (s *RetrieveMemoryRequestQuery) SetNamespace(v string) *RetrieveMemoryRequestQuery {
	s.Namespace = &v
	return s
}

func (s *RetrieveMemoryRequestQuery) SetUserId(v string) *RetrieveMemoryRequestQuery {
	s.UserId = &v
	return s
}

func (s *RetrieveMemoryRequestQuery) Validate() error {
	return dara.Validate(s)
}
