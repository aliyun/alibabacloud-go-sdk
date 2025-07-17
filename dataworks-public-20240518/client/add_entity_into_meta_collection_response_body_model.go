// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntityIntoMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddEntityIntoMetaCollectionResponseBody
	GetRequestId() *string
}

type AddEntityIntoMetaCollectionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEntityIntoMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEntityIntoMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntityIntoMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEntityIntoMetaCollectionResponseBody) SetRequestId(v string) *AddEntityIntoMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEntityIntoMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
