// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateDomainMetaResponseBody
	GetData() *string
	SetRequestId(v string) *CreateDomainMetaResponseBody
	GetRequestId() *string
}

type CreateDomainMetaResponseBody struct {
	// The ID of the newly created list.
	//
	// example:
	//
	// ladl-6f1exxxxx6ab59
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C37AE32-A5C0-5E0F-9EC3-399B83102ED1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainMetaResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateDomainMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainMetaResponseBody) SetData(v string) *CreateDomainMetaResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDomainMetaResponseBody) SetRequestId(v string) *CreateDomainMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
