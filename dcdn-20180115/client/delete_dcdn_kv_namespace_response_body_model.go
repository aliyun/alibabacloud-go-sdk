// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnKvNamespaceResponseBody
	GetRequestId() *string
}

type DeleteDcdnKvNamespaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnKvNamespaceResponseBody) SetRequestId(v string) *DeleteDcdnKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
