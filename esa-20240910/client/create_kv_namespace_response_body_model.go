// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKvNamespaceResponseBody
	GetDescription() *string
	SetNamespace(v string) *CreateKvNamespaceResponseBody
	GetNamespace() *string
	SetNamespaceId(v string) *CreateKvNamespaceResponseBody
	GetNamespaceId() *string
	SetRequestId(v string) *CreateKvNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateKvNamespaceResponseBody
	GetStatus() *string
}

type CreateKvNamespaceResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateKvNamespaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateKvNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKvNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateKvNamespaceResponseBody) SetDescription(v string) *CreateKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetNamespace(v string) *CreateKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetNamespaceId(v string) *CreateKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetRequestId(v string) *CreateKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) SetStatus(v string) *CreateKvNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
