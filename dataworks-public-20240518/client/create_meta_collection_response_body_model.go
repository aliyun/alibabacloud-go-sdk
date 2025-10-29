// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateMetaCollectionResponseBody
	GetId() *string
	SetRequestId(v string) *CreateMetaCollectionResponseBody
	GetRequestId() *string
}

type CreateMetaCollectionResponseBody struct {
	// The collection ID returned after a successful creation.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaCollectionResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetaCollectionResponseBody) SetId(v string) *CreateMetaCollectionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetRequestId(v string) *CreateMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
