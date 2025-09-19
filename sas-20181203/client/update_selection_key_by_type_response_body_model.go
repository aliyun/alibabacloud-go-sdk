// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSelectionKeyByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSelectionKeyByTypeResponseBody
	GetRequestId() *string
}

type UpdateSelectionKeyByTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DBABBC2E-26DF-5586-BF7C-4FC846EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSelectionKeyByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSelectionKeyByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSelectionKeyByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSelectionKeyByTypeResponseBody) SetRequestId(v string) *UpdateSelectionKeyByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSelectionKeyByTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
