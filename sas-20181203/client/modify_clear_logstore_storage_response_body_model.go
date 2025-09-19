// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClearLogstoreStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClearLogstoreStorageResponseBody
	GetRequestId() *string
}

type ModifyClearLogstoreStorageResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DC84C453-8561-5EC4-B0E9-44E67ACCB5B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClearLogstoreStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClearLogstoreStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClearLogstoreStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClearLogstoreStorageResponseBody) SetRequestId(v string) *ModifyClearLogstoreStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClearLogstoreStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
