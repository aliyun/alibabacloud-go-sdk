// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMetaCollectionResponseBody
	GetRequestId() *string
}

type DeleteMetaCollectionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 45D14A7A-7C28-5547-AB0A-35FBCD9DE7B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaCollectionResponseBody) SetRequestId(v string) *DeleteMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
