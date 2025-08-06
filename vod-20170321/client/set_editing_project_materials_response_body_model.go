// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type SetEditingProjectMaterialsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 746FFA07-8BBB-46*****B1-3E94E3B2915E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetEditingProjectMaterialsResponseBody) SetRequestId(v string) *SetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEditingProjectMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}
