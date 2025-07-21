// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeyDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKeyDescriptionResponseBody
	GetRequestId() *string
}

type UpdateKeyDescriptionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKeyDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeyDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKeyDescriptionResponseBody) SetRequestId(v string) *UpdateKeyDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKeyDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
