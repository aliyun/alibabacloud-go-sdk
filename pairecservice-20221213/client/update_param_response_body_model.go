// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateParamResponseBody
	GetRequestId() *string
}

type UpdateParamResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BBD41FBF-E75C-551A-92FA-CAD654AA006F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateParamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateParamResponseBody) SetRequestId(v string) *UpdateParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateParamResponseBody) Validate() error {
	return dara.Validate(s)
}
