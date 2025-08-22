// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnKvResponseBody
	GetRequestId() *string
}

type DeleteDcdnKvResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnKvResponseBody) SetRequestId(v string) *DeleteDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}
