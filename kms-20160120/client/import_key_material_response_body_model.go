// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyMaterialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ImportKeyMaterialResponseBody
	GetRequestId() *string
}

type ImportKeyMaterialResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ec1017cf-ead4-f3ca-babc-c3b34f3dbecb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyMaterialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKeyMaterialResponseBody) SetRequestId(v string) *ImportKeyMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKeyMaterialResponseBody) Validate() error {
	return dara.Validate(s)
}
