// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody
	GetData() *CreateNamespaceResponseBodyData
	SetRequestId(v string) *CreateNamespaceResponseBody
	GetRequestId() *string
}

type CreateNamespaceResponseBody struct {
	// The namespace creation result.
	Data *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) GetData() *CreateNamespaceResponseBodyData {
	return s.Data
}

func (s *CreateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNamespaceResponseBodyData struct {
	// The namespace ID. The default namespace cannot be queried, modified, or deleted.
	//
	// example:
	//
	// a2a9310a-9d91-4283-b4e2-844f6d45fe64
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceId(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
