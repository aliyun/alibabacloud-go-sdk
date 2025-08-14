// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type DeleteEditingProjectMaterialsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******36-3C1E-4417-BDB2-1E034F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEditingProjectMaterialsResponseBody) SetRequestId(v string) *DeleteEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}
