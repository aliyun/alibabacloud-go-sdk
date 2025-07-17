// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaCollectionResponseBody
	GetRequestId() *string
}

type UpdateMetaCollectionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0E1C0122-F79F-5C26-B546-47A321691868
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaCollectionResponseBody) SetRequestId(v string) *UpdateMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
