// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKeyMaterialResponseBody
	GetRequestId() *string
}

type DeleteKeyMaterialResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKeyMaterialResponseBody) SetRequestId(v string) *DeleteKeyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeyMaterialResponseBody) Validate() error {
	return dara.Validate(s)
}
